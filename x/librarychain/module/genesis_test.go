package librarychain_test

import (
	"testing"

	keepertest "librarychain/testutil/keeper"
	"librarychain/testutil/nullify"
	librarychain "librarychain/x/librarychain/module"
	"librarychain/x/librarychain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BookList: []types.Book{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BookCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LibrarychainKeeper(t)
	librarychain.InitGenesis(ctx, k, genesisState)
	got := librarychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BookList, got.BookList)
	require.Equal(t, genesisState.BookCount, got.BookCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
