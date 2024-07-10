package keeper

import (
	"context"

	"cosa/x/cosa/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CloseAuction(goCtx context.Context, msg *types.MsgCloseAuction) (*types.MsgCloseAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found:= k.GetAuction(ctx, msg.Id)

	if !found {
		return nil, sdkerrors.Wrapf(sdkerror.ErrKeyNotFound, "auction %d doesn't exist", msg.Id)
	}

	if auction.Status != Approved {
		return nil, sdkerrors.Wrap(sdkerror.ErrUnauthorized, "auction is not approved")
	}

	if ctx.BlockTime().Before(auction.Endtime.AsTime()) {
        return nil, sdkerrors.Wrap(sdkerror.ErrUnauthorized, "auction has not ended yet")
    }

	if len(auction.Bids) == 0 {
        return nil, sdkerrors.Wrap(sdkerror.ErrUnauthorized, "no bids placed")
    }

	highestBid:= auction.HighestBid
	winner:= auction.HighestBidder
	
	auction.Status = Closed
	auction.Owner = winner
	auction.SalePrice = highestBid

	k.SetAuction(ctx, auction)

	return &types.MsgCloseAuctionResponse{
		Winner: winner,
		HighestBid: highestBid,
	}, nil
}
