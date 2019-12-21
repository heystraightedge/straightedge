package main

import (
	"bufio"
	"io"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"

	"github.com/heystraightedge/straightedge/sr25519"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"

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
	cmd.AddCommand(
		clientkeys.MnemonicKeyCommand(),
		addCmd,
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
		return keys.NewInMemory(keys.WithKeygenFunc()), nil
	}

	return clientkeys.NewKeyringFromHomeFlag(buf,
		keys.WithKeygenFunc(straightedgeKeygenFunc),
		keys.WithDeriveFunction(straightedgeKeygenFunc),
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

// Straightedge KeyGenFunc currently supports secp256k1 and sr25119 keys
func straightedgeKeygenFunc(bz []byte, algo keys.SigningAlgo) tmcrypto.PrivKey {
	if algo == keys.Secp256k1 {
		var bzArr [32]byte
		copy(bzArr[:], bz)
		return secp256k1.PrivKeySecp256k1(bzArr)
	} else if algo == keys.Sr25519 {
		var bzArr [32]byte
		copy(bzArr[:], bz)
		return sr25519.PrivKeySr25519(bzArr)
	}
	return nil
}

// Straightedge DeriveFunc currently supports secp256k1 and sr25119 keys
func straightedgeDeriveFunc(mnemonic string, bip39Passphrase, hdPath string, algo SigningAlgo) ([]byte, error) {
	if algo == keys.Secp256k1 {
		return keys.baseDeriveKey(mnemonic, bip39Passphrase, hdPath, algo)
	} else if algo == keys.Sr25519 {

	}
	return nil, keys.ErrUnsupportedSigningAlgo
}
