package types

// NewGenesisState creates a new genesis state.
func NewGenesisState(disabledRoutes []string) *GenesisState {
	return &GenesisState{DisabledRoutes: disabledRoutes}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *GenesisState { return NewGenesisState([]string{}) }
