package app

import (
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestExport(t *testing.T) {
	db := db.NewMemDB()
	// logger, db, traceStore, true, wasm.EnableAllProposals, skipUpgradeHeights,
	app := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, wasm.EnableAllProposals, nil, "", 0, MakeEncodingConfig(), simapp.EmptyAppOptions{})

	setGenesis(app)

	// Making a new app object with the db, so that initchain hasn't been called
	newApp := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, wasm.EnableAllProposals, nil, "", 0, MakeEncodingConfig(), simapp.EmptyAppOptions{})
	_, err := newApp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlockedAddrs(t *testing.T) {
	db := db.NewMemDB()
	app := NewStraightedgeApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, wasm.EnableAllProposals, nil, "", 0, MakeEncodingConfig(), simapp.EmptyAppOptions{})

	for acc := range maccPerms {
		require.Equal(t, !allowedReceivingModAcc[acc], app.BankKeeper.BlockedAddr(app.AccountKeeper.GetModuleAddress(acc)),
			"ensure that blocked addresses are properly set in bank keeper")
	}
}

func setGenesis(app *StraightedgeApp) error {
	genesisState := NewDefaultGenesisState(app.appCodec)

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
