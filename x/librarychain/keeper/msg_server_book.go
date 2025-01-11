package keeper

import (
	"context"
	"fmt"

	"librarychain/x/librarychain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBook(goCtx context.Context, msg *types.MsgCreateBook) (*types.MsgCreateBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var book = types.Book{
		Creator: msg.Creator,
		Title:   msg.Title,
		Author:  msg.Author,
	}

	id := k.AppendBook(
		ctx,
		book,
	)

	return &types.MsgCreateBookResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBook(goCtx context.Context, msg *types.MsgUpdateBook) (*types.MsgUpdateBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var book = types.Book{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Author:  msg.Author,
	}

	// Checks that the element exists
	val, found := k.GetBook(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBook(ctx, book)

	return &types.MsgUpdateBookResponse{}, nil
}

func (k msgServer) DeleteBook(goCtx context.Context, msg *types.MsgDeleteBook) (*types.MsgDeleteBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetBook(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBook(ctx, msg.Id)

	return &types.MsgDeleteBookResponse{}, nil
}
