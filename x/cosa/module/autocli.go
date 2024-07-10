package cosa

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "cosa/api/cosa/cosa"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateAuction",
					Use:            "create-auction [item] [starting-price] [duration] [status]",
					Short:          "Send a createAuction tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "item"}, {ProtoField: "startingPrice"}, {ProtoField: "duration"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "ApproveAuction",
					Use:            "approve-auction [item] [id] [starting-price]",
					Short:          "Send a approveAuction tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "item"}, {ProtoField: "id"}, {ProtoField: "startingPrice"}},
				},
				{
					RpcMethod:      "CreatBid",
					Use:            "creat-bid [item] [auction-id] [bid-amount] [bidder]",
					Short:          "Send a creatBid tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "item"}, {ProtoField: "auctionId"}, {ProtoField: "bidAmount"}, {ProtoField: "bidder"}},
				},
				{
					RpcMethod:      "CloseAuction",
					Use:            "close-auction [id] [highest-bid]",
					Short:          "Send a closeAuction tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "highestBid"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
