package app

import (
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestExport(t *testing.T) {
	db := db.NewMemDB()
	app := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, wasm.EnableAllProposals, nil)
	setGenesis(app)

	// Making a new app object with the db, so that initchain hasn't been called
	newApp := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, wasm.EnableAllProposals, nil)
	_, _, err := newApp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := db.NewMemDB()
	app := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0, wasm.EnableAllProposals, nil)

	for acc := range maccPerms {
		require.True(t, app.BankKeeper.BlacklistedAddr(app.supplyKeeper.GetModuleAddress(acc)))
	}
}

func setGenesis(app *StraightedgeApp) error {
	genesisState := NewDefaultGenesisState()

	stateBytes, err := codec.MarshalJSONIndent(app.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	return nil
}
