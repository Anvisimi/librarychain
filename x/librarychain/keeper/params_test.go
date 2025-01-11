package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "librarychain/testutil/keeper"
	"librarychain/x/librarychain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LibrarychainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
