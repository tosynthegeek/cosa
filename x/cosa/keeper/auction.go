package keeper

import (
	"cosa/x/cosa/types"
	"encoding/binary"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) int {
	count:= k.GetAuctionCount(ctx)
	auction.Id = uint64(count)
	storeAdapter:= runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store:= prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	idBytes:= IDBytes(auction.Id)
	auctionBytes:= k.cdc.MustMarshal(&auction)

	store.Set(idBytes, auctionBytes)

	return count
}

func (k Keeper) GetAuction(ctx sdk.Context, id uint64) (types.Auction, bool) {
	idBytes:= IDBytes(id)
	storeAdapter:= runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store:= prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))
	auctionBytes:= store.Get(idBytes)
	if auctionBytes == nil {
		return types.Auction{}, false
	}
	var auction types.Auction
	k.cdc.Unmarshal(auctionBytes, &auction)

	return auction, true
}

func (k Keeper) UpdateAuction(ctx sdk.Context, auction types.Auction) error {
	existingAuction, found:= k.GetAuction(ctx, auction.Id)
	if !found {
		return sdkerrors.Wrapf(sdkerror.ErrKeyNotFound, "auction %d doesnt exist", auction.Id)
	}

	// Optional: Perform any necessary validations
	// For example, you might want to prevent updates to certain fields
	if auction.Creator != existingAuction.Creator {
		return sdkerrors.Wrap(sdkerror.ErrUnauthorized, "cannot change bid creator")
	}

	k.SetAuction(ctx, auction)

	return nil
}

func (k Keeper) GetAllAuctions(ctx sdk.Context) []types.Auction {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var auctions []types.Auction
	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		k.cdc.Unmarshal(iterator.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
} 

func (k Keeper) GetAuctionCount(ctx sdk.Context) int {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	count := 0
	for ; iterator.Valid(); iterator.Next() {
		count++
	}

	return count
}

func IDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
