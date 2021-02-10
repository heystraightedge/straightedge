package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyDisabledRoutes, []string{}, ValidateDisabledRoutes),
	)
}

// ValidateDisabledRoutes validate routes
func ValidateDisabledRoutes(i interface{}) error {
	routes, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, route := range routes {
		if !IsAlphaNumeric(route) {
			return fmt.Errorf("invalid route: %s", route)
		}
	}

	return nil
}

// GetDisabledRoutes returns disabled routes
func (rtr *Router) GetDisabledRoutes(ctx sdk.Context) []string {
	var disabledRoutes []string
	rtr.paramSpace.Get(ctx, ParamStoreKeyDisabledRoutes, &disabledRoutes)
	return disabledRoutes
}

// SetDisabledRoutes set disabled routes
func (rtr *Router) SetDisabledRoutes(ctx sdk.Context, disabledRoutes []string) {
	rtr.paramSpace.Set(ctx, ParamStoreKeyDisabledRoutes, disabledRoutes)
}
