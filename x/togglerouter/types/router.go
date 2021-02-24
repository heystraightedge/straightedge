package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// ModuleName for router
	ModuleName = "router"
)

// Router describes custom router
type Router struct {
	// The reference to the Paramstore to get and set gov specific params
	paramSpace paramtypes.Subspace
	routes     map[string]sdk.Handler
}

var _ sdk.Router = NewRouter(paramtypes.Subspace{})

// NewRouter returns a reference to a new router.
func NewRouter(paramSpace paramtypes.Subspace) *Router {
	return &Router{
		paramSpace: paramSpace,
		routes:     make(map[string]sdk.Handler),
	}
}

// AddRoute adds a route path to the router with a given handler. The route must
// be alphanumeric.
func (rtr *Router) AddRoute(r sdk.Route) sdk.Router {
	if !types.IsAlphaNumeric(r.Path()) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.routes[r.Path()] != nil {
		panic(fmt.Sprintf("route %s has already been initialized", r.Path()))
	}

	rtr.routes[r.Path()] = r.Handler()
	return rtr
}

// Route returns a handler for a given route path.
//
// TODO: Handle expressive matches.
func (rtr *Router) Route(ctx sdk.Context, path string) sdk.Handler {
	if rtr.GetRouteDisabled(ctx, path) {
		return nil
	}

	return rtr.routes[path]
}

// GetRouteDisabled returns whether a specific route is disabled from the global param store
func (rtr *Router) GetRouteDisabled(ctx sdk.Context, route string) bool {
	var disabledRoutes []string
	rtr.paramSpace.Get(ctx, ParamStoreKeyDisabledRoutes, &disabledRoutes)
	for _, disabledRoute := range disabledRoutes {
		if route == disabledRoute {
			return true
		}
	}
	return false
}

// SetRouteDisabled sets whether a specific route is disabled from the global param store
func (rtr *Router) SetRouteDisabled(ctx sdk.Context, route string, disabled bool) {
	if disabled {
		disabledRoutes := rtr.GetDisabledRoutes(ctx)
		for _, disabledRoute := range disabledRoutes {
			if route == disabledRoute {
				return
			}
		}
		disabledRoutes = append(disabledRoutes, route)
		rtr.SetDisabledRoutes(ctx, disabledRoutes)
	} else {
		newDisabled := []string{}
		disabledRoutes := rtr.GetDisabledRoutes(ctx)
		for _, disabledRoute := range disabledRoutes {
			if route == disabledRoute {
				newDisabled = append(newDisabled, route)
			}
		}
		rtr.SetDisabledRoutes(ctx, newDisabled)
	}
}
