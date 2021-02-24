package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState represents the genesis state of the blockchain. It is a map from module names to module genesis states.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONMarshaler) GenesisState {
	return ModuleBasics.DefaultGenesis(cdc)
}
