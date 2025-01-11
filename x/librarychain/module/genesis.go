package librarychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"librarychain/x/librarychain/keeper"
	"librarychain/x/librarychain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the book
	for _, elem := range genState.BookList {
		k.SetBook(ctx, elem)
	}

	// Set book count
	k.SetBookCount(ctx, genState.BookCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.BookList = k.GetAllBook(ctx)
	genesis.BookCount = k.GetBookCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
