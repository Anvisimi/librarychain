package keeper_test

import (
	"context"
	"testing"

	keepertest "librarychain/testutil/keeper"
	"librarychain/testutil/nullify"
	"librarychain/x/librarychain/keeper"
	"librarychain/x/librarychain/types"

	"github.com/stretchr/testify/require"
)

func createNBook(keeper keeper.Keeper, ctx context.Context, n int) []types.Book {
	items := make([]types.Book, n)
	for i := range items {
		items[i].Id = keeper.AppendBook(ctx, items[i])
	}
	return items
}

func TestBookGet(t *testing.T) {
	keeper, ctx := keepertest.LibrarychainKeeper(t)
	items := createNBook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBook(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBookRemove(t *testing.T) {
	keeper, ctx := keepertest.LibrarychainKeeper(t)
	items := createNBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBook(ctx, item.Id)
		_, found := keeper.GetBook(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.LibrarychainKeeper(t)
	items := createNBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBook(ctx)),
	)
}

func TestBookCount(t *testing.T) {
	keeper, ctx := keepertest.LibrarychainKeeper(t)
	items := createNBook(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBookCount(ctx))
}
