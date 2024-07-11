package keeper

import (
	"context"
	"fmt"
	"time"

	"cosa/x/cosa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	endTime := time.Now().Add(time.Duration(msg.Duration))
	fmt.Printf("endtime: %v\n", endTime)
	endtime:= sdk.FormatTimeString(endTime)

	auction := types.Auction{
		Item:          msg.Item,
		Creator:       msg.Creator,
		StartingPrice: msg.StartingPrice,
		Duration:      msg.Duration,
		Endtime:       endtime,
		Status:        Pending,
	}

	k.SetAuction(ctx, auction)

	return &types.MsgCreateAuctionResponse{
		Id: auction.Id,
	}, nil
}
