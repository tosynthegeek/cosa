package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCloseAuction{}

func NewMsgCloseAuction(creator string, id uint64, highestBid uint64) *MsgCloseAuction {
	return &MsgCloseAuction{
		Creator:    creator,
		Id:         id,
		HighestBid: highestBid,
	}
}

func (msg *MsgCloseAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
