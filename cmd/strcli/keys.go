package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"

	substratebip39 "github.com/sikkatech/go-substrate-bip39"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/sr25519"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagDryRun = "dry-run"
)

// keyCommands registers a sub-tree of commands to interact with
// local private key storage.
func keyCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
		Long: `Keys allows you to manage your local keystore for tendermint.
    These keys may be in any format supported by go-crypto and can be
    used by light-clients, full nodes, or any other application that
    needs to sign with a private key.`,
	}
	addCmd := clientkeys.AddKeyCommand()
	addCmd.RunE = runAddCmd
	migrateCmd := MigrateKeyCommand()
	cmd.AddCommand(
		clientkeys.MnemonicKeyCommand(),
		addCmd,
		migrateCmd,
		clientkeys.ExportKeyCommand(),
		clientkeys.ImportKeyCommand(),
		clientkeys.ListKeysCmd(),
		clientkeys.ShowKeysCmd(),
		flags.LineBreak,
		clientkeys.DeleteKeyCommand(),
		clientkeys.UpdateKeyCommand(),
		clientkeys.ParseKeyStringCommand(),
		clientkeys.MigrateCommand(),
		flags.LineBreak,
	)
	return cmd
}

func getKeybase(cmd *cobra.Command, dryrun bool, buf io.Reader) (keys.Keybase, error) {
	if dryrun {
		return keys.NewInMemory(
			keys.WithKeygenFunc(straightedgeKeygenFunc),
			keys.WithDeriveFunc(straightedgeDeriveFunc),
			keys.WithSupportedAlgos([]keys.SigningAlgo{keys.Secp256k1, keys.Sr25519}),
			keys.WithSupportedAlgosLedger([]keys.SigningAlgo{keys.Secp256k1}),
		), nil
	}

	return keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), buf,
		keys.WithKeygenFunc(straightedgeKeygenFunc),
		keys.WithDeriveFunc(straightedgeDeriveFunc),
		keys.WithSupportedAlgos([]keys.SigningAlgo{keys.Secp256k1, keys.Sr25519}),
		keys.WithSupportedAlgosLedger([]keys.SigningAlgo{keys.Secp256k1}),
	)
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	kb, err := getKeybase(cmd, viper.GetBool(flagDryRun), inBuf)
	if err != nil {
		return err
	}

	return clientkeys.RunAddCmd(cmd, args, kb, inBuf)
}

func MigrateKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate <name>",
		Short: "Migrate balance from an existing sr25519 key to another locally stored key of the given name",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		RunE:  runMigrateCmd,
	}

	cmd.SetOut(cmd.OutOrStdout())
	cmd.SetErr(cmd.ErrOrStderr())

	return cmd
}

// MigrateCmd asks you for your mnemonic, then
// derives your sr25519 key and a secp256k1 key
func runMigrateCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	kb, err := getKeybase(cmd, viper.GetBool(flagDryRun), inBuf)
	if err != nil {
		return err
	}

	name := args[0]

	// bip39 passpharse is ""
	var mnemonic, bip39Passphrase string

	mnemonic, err = input.GetString("Enter your sr25519 key mnemonic", inBuf)
	if err != nil {
		return err
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		return errors.New("invalid mnemonic")
	}

	sr25519HdPath := ""
	srPrivKeyBz, err := straightedgeDeriveFunc(mnemonic, bip39Passphrase, sr25519HdPath, keys.Sr25519)
	if err != nil {
		return err
	}
	_, err = straightedgeKeygenFunc(srPrivKeyBz, keys.Sr25519)
	if err != nil {
		return err
	}

	info, err := kb.Get(name)
	if err != nil {
		return err
	}
	payeeAddr := info.GetAddress()

	fmt.Println("Please confirm you want to send funds to address:\n", payeeAddr.String())
	yesno, err := input.GetString("Enter y/n", inBuf)
	if err != nil {
		return err
	}
	if yesno != "y" && yesno != "yes" {
		return nil
	}

	return nil
}

// Straightedge KeyGenFunc currently supports secp256k1 and sr25119 keys
func straightedgeKeygenFunc(bz []byte, algo keys.SigningAlgo) (tmcrypto.PrivKey, error) {
	if algo == keys.Secp256k1 {
		return keys.SecpPrivKeyGen(bz), nil
	} else if algo == keys.Sr25519 {
		var bzArr [32]byte
		copy(bzArr[:], bz)

		privKey := sr25519.PrivKeySr25519(bzArr)
		fmt.Println("Derived Sr25519 Pubkey: ")
		fmt.Println(privKey.PubKey())

		return privKey, nil
	}
	return nil, keys.ErrUnsupportedSigningAlgo
}

// Straightedge DeriveFunc currently supports secp256k1 and sr25119 keys
func straightedgeDeriveFunc(mnemonic string, bip39Passphrase, hdPath string, algo keys.SigningAlgo) ([]byte, error) {
	if algo == keys.Secp256k1 {
		return keys.StdDeriveKey(mnemonic, bip39Passphrase, hdPath, algo)
	} else if algo == keys.Sr25519 {
		return sr25519DeriveFunction(mnemonic, bip39Passphrase, hdPath, algo)
	}
	return nil, keys.ErrUnsupportedSigningAlgo
}

func sr25519DeriveFunction(mnemonic string, bip39Passphrase, hdPath string, algo keys.SigningAlgo) ([]byte, error) {
	if algo == keys.Sr25519 {
		seed, err := substratebip39.SeedFromMnemonic(mnemonic, bip39Passphrase)
		return seed[:], err
	}
	return nil, keys.ErrUnsupportedSigningAlgo
}
