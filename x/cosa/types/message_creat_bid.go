package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatBid{}

func NewMsgCreatBid(creator string, item string, auctionId uint64, bidAmount uint64, bidder string) *MsgCreatBid {
	return &MsgCreatBid{
		Creator:   creator,
		Item:      item,
		AuctionId: auctionId,
		BidAmount: bidAmount,
		Bidder:    bidder,
	}
}

func (msg *MsgCreatBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
