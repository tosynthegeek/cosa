package keeper

import (
	"context"

	"cosa/x/cosa/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatBid(goCtx context.Context, msg *types.MsgCreatBid) (*types.MsgCreatBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerror.ErrKeyNotFound, "auction %d doesnt exist", auction.Id)
	}

	if auction.Status != Approved {
		return nil, sdkerrors.Wrapf(sdkerror.ErrKeyNotFound, "auction %d doesnt is not approved", auction.Id)
	}

	if ctx.BlockTime().After(auction.Endtime.AsTime()) {
		return nil, sdkerrors.Wrap(sdkerror.ErrUnauthorized, "auction has ended")
	}

	bid := types.Bid{
		Bidder: msg.Bidder,
		Amount: msg.BidAmount,
	}

	auction.Bids = append(auction.Bids, &bid)

	if msg.BidAmount > auction.HighestBid {
		auction.HighestBid = msg.BidAmount
		auction.HighestBidder = msg.Bidder
	}

	k.SetAuction(ctx, auction)

	return &types.MsgCreatBidResponse{
		Id: msg.AuctionId,
	}, nil
}
