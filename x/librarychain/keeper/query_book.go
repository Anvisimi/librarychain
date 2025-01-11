package keeper

import (
	"context"

	"librarychain/x/librarychain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BookAll(ctx context.Context, req *types.QueryAllBookRequest) (*types.QueryAllBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var books []types.Book

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bookStore := prefix.NewStore(store, types.KeyPrefix(types.BookKey))

	pageRes, err := query.Paginate(bookStore, req.Pagination, func(key []byte, value []byte) error {
		var book types.Book
		if err := k.cdc.Unmarshal(value, &book); err != nil {
			return err
		}

		books = append(books, book)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBookResponse{Book: books, Pagination: pageRes}, nil
}

func (k Keeper) Book(ctx context.Context, req *types.QueryGetBookRequest) (*types.QueryGetBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	book, found := k.GetBook(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBookResponse{Book: book}, nil
}
