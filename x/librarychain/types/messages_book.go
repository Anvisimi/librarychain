package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBook{}

func NewMsgCreateBook(creator string, title string, author string) *MsgCreateBook {
	return &MsgCreateBook{
		Creator: creator,
		Title:   title,
		Author:  author,
	}
}

func (msg *MsgCreateBook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBook{}

func NewMsgUpdateBook(creator string, id uint64, title string, author string) *MsgUpdateBook {
	return &MsgUpdateBook{
		Id:      id,
		Creator: creator,
		Title:   title,
		Author:  author,
	}
}

func (msg *MsgUpdateBook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBook{}

func NewMsgDeleteBook(creator string, id uint64) *MsgDeleteBook {
	return &MsgDeleteBook{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteBook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
