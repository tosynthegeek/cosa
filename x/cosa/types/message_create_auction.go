package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAuction{}

func NewMsgCreateAuction(creator string, item string, startingPrice uint64, duration uint64, status string) *MsgCreateAuction {
	return &MsgCreateAuction{
		Creator:       creator,
		Item:          item,
		StartingPrice: startingPrice,
		Duration:      duration,
		Status:        status,
	}
}

func (msg *MsgCreateAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
