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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LibrarychainKeeper(t)
	librarychain.InitGenesis(ctx, k, genesisState)
	got := librarychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
