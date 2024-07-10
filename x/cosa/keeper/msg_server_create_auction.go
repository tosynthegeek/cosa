package keeper

import (
	"context"

	"cosa/x/cosa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var Pending = "Pending"
var Approved = "Approved"
var Closed = "Closed"

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction:= types.Auction {
		Item: msg.Item,
		Creator: msg.Creator,
		StartingPrice: msg.StartingPrice,
		Duration: msg.Duration,
		Status: Pending,
	}

	k.SetAuction(ctx, auction)

	return &types.MsgCreateAuctionResponse{
		Id: auction.Id,
	}, nil
}
