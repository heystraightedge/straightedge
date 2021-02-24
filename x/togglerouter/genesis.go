package togglerouter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/heystraightedge/straightedge/x/togglerouter/types"
)

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data types.GenesisState) error {
	return types.ValidateDisabledRoutes(data.DisabledRoutes)
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, router *types.Router, data types.GenesisState) {
	router.SetDisabledRoutes(ctx, data.DisabledRoutes)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, router *types.Router) types.GenesisState {
	return *types.NewGenesisState(router.GetDisabledRoutes(ctx))
}
