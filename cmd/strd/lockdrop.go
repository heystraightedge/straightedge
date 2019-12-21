package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/heystraightedge/straightedge/sr25519"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/genutil"
)

type GenBalances struct {
	GenBalances [][]string `json:"balances"`
}

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func ImportLockdropBalancesCmd(
	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "import-lockdrop-balances [denom] [file]",
		Short: "Import balances from lockdrop to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			denom := args[0]
			filepath := args[1]

			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			authGenState := auth.GetGenesisStateFromAppState(cdc, appState)

			jsonFile, err := os.Open(filepath)
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			var genBalances GenBalances
			json.Unmarshal(byteValue, &genBalances)
			for i := 0; i < len(genBalances.GenBalances); i++ {
				fmt.Println("PubKey: " + genBalances.GenBalances[i][0])
				fmt.Println("Amount: " + genBalances.GenBalances[i][1])

				bz, err := hex.DecodeString(genBalances.GenBalances[i][0])
				if err != nil {
					return err
				}
				var bzArr [32]byte
				copy(bzArr[:], bz[:])

				pubKey := sr25519.PubKeySr25519(bzArr)

				addr := sdk.AccAddress(pubKey.Address().Bytes())

				amount, ok := sdk.NewIntFromString(genBalances.GenBalances[i][1])
				if !ok {
					return errors.New("couldn't parse balance")
				}

				coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

				// create concrete account type based on input parameters
				var genAccount authexported.GenesisAccount
				baseAccount := auth.NewBaseAccount(addr, coins.Sort(), nil, 0, 0)
				genAccount = baseAccount

				if err := genAccount.Validate(); err != nil {
					return fmt.Errorf("failed to validate new genesis account: %w", err)
				}

				if authGenState.Accounts.Contains(addr) {
					return fmt.Errorf("cannot add account at existing address %s", addr)
				}

				// Add the new account to the set of genesis accounts and sanitize the
				// accounts afterwards.
				authGenState.Accounts = append(authGenState.Accounts, genAccount)
				authGenState.Accounts = auth.SanitizeGenesisAccounts(authGenState.Accounts)
			}

			authGenStateBz, err := cdc.MarshalJSON(authGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal auth genesis state: %w", err)
			}

			appState[auth.ModuleName] = authGenStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")

	return cmd
}
