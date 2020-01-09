package togglerouter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState is the bank state that must be provided at genesis.
type GenesisState struct {
	DisabledRoutes []string `json:"disabled_routes"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(disabledRoutes []string) GenesisState {
	return GenesisState{DisabledRoutes: disabledRoutes}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState { return NewGenesisState([]string{}) }

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return validateDisabledRoutes(data.DisabledRoutes)
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, router *Router, data GenesisState) {
	router.setDisabledRoutes(ctx, data.DisabledRoutes)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, router *Router) GenesisState {
	return NewGenesisState(router.getDisabledRoutes(ctx))
}
