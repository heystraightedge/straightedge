package togglerouter

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	params "github.com/cosmos/cosmos-sdk/x/params/subspace"
)

const (

	// DefaultParamspace for params keeper
	DefaultParamspace = ModuleName
)

var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

type Router struct {
	// The reference to the Paramstore to get and set gov specific params
	paramSpace types.ParamSubspace

	// Reserved codespace
	codespace sdk.CodespaceType

	paramTable params.KeyTable

	routeEnabled map[string]bool

	routes map[string]sdk.Handler
}

var _ sdk.Router = NewRouter()

// NewRouter returns a reference to a new router.
func NewRouter(paramSpace types.ParamSubspace, codespace sdk.CodespaceType) *Router {
	return &Router{
		paramSpace: paramSpace,
		codespace:  codespace,
		paramTable: params.NewKeyTable(),
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

	rtr.addRouteParamKeyTable(path)

	rtr.routes[path] = h
	return rtr
}

// Route returns a handler for a given route path.
//
// TODO: Handle expressive matches.
func (rtr *Router) Route(path string) sdk.Handler {
	return rtr.routes[path]
}

// GetRouteEnabled returns whether a specific route is enabled from the global param store
func (rtr *Router) GetRouteEnabled(ctx sdk.Context, route string) bool {
	var routeEnabled bool
	rtr.paramSpace.Get(ctx, []byte(route), &routeEnabled)
	return routeEnabled
}

// SetRouteEnabled sets whether a specific route is enabled from the global param store
func (rtr *Router) SetRouteEnabled(ctx sdk.Context, route string, enabled bool) {
	rtr.paramSpace.Set(ctx, []byte(route), &enabled)
}

func (rtr *Router) addRouteParamKeyTable(path string) {
	rtr.paramTable.RegisterType(params.NewParamSetPair([]byte(path), true,
		func(i interface{}) error {
			_, ok := i.(bool)
			if !ok {
				return fmt.Errorf("invalid parameter type: %T", i)
			}

			return nil
		},
	))
}

// ParamKeyTable type declaration for parameters
func (rtr *Router) ParamKeyTable() params.KeyTable {
	return rtr.paramTable
}
