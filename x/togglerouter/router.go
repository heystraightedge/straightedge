package togglerouter

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	ModuleName = "router"

	// DefaultParamspace for params keeper
	DefaultParamspace = ModuleName
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

type Router struct {
	// The reference to the Paramstore to get and set gov specific params
	paramSpace params.Subspace
	routes     map[string]sdk.Handler
}

var _ sdk.Router = NewRouter(params.Subspace{})

// NewRouter returns a reference to a new router.
func NewRouter(paramSpace params.Subspace) *Router {
	return &Router{
		paramSpace: paramSpace,
		routes:     make(map[string]sdk.Handler),
	}
}

// AddRoute adds a route path to the router with a given handler. The route must
// be alphanumeric.
func (rtr *Router) AddRoute(path string, h sdk.Handler) sdk.Router {
	if !isAlphaNumeric(path) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.routes[path] != nil {
		panic(fmt.Sprintf("route %s has already been initialized", path))
	}

	rtr.routes[path] = h
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

// GetRouteEnabled returns whether a specific route is disabled from the global param store
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

// SetRouteEnabled sets whether a specific route is disabled from the global param store
func (rtr *Router) SetRouteDisabled(ctx sdk.Context, route string, disabled bool) {
	if disabled {
		disabledRoutes := rtr.getDisabledRoutes(ctx)
		for _, disabledRoute := range disabledRoutes {
			if route == disabledRoute {
				return
			}
		}
		disabledRoutes = append(disabledRoutes, route)
		rtr.setDisabledRoutes(ctx, disabledRoutes)
	} else {
		newDisabled := []string{}
		disabledRoutes := rtr.getDisabledRoutes(ctx)
		for _, disabledRoute := range disabledRoutes {
			if route == disabledRoute {
				newDisabled = append(newDisabled, route)
			}
		}
		rtr.setDisabledRoutes(ctx, newDisabled)
	}
}
