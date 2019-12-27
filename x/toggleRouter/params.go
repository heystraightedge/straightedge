package togglerouter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// ParamStoreKeyDisabledRoutes is store's key for DisabledRoutes
var ParamStoreKeyDisabledRoutes = []byte("disabledroutes")

// ParamKeyTable type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		params.NewParamSetPair(ParamStoreKeyDisabledRoutes, []string{}, validateDisabledRoutes),
	)
}

func validateDisabledRoutes(i interface{}) error {
	routes, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, route := range routes {
		if !isAlphaNumeric(route) {
			return fmt.Errorf("invalid route: %s", route)
		}
	}

	return nil
}

func (rtr *Router) getDisabledRoutes(ctx sdk.Context) []string {
	var disabledRoutes []string
	rtr.paramSpace.Get(ctx, ParamStoreKeyDisabledRoutes, &disabledRoutes)
	return disabledRoutes
}

func (rtr *Router) setDisabledRoutes(ctx sdk.Context, disabledRoutes []string) {
	rtr.paramSpace.Set(ctx, ParamStoreKeyDisabledRoutes, disabledRoutes)
}
