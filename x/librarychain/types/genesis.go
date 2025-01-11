package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		BookList: []Book{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in book
	bookIdMap := make(map[uint64]bool)
	bookCount := gs.GetBookCount()
	for _, elem := range gs.BookList {
		if _, ok := bookIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for book")
		}
		if elem.Id >= bookCount {
			return fmt.Errorf("book id should be lower or equal than the last id")
		}
		bookIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
