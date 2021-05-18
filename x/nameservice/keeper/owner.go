package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var ownersStoreKey = sdk.NewKVStoreKey("temp")

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	store := ctx.KVStore(ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	store := ctx.KVStore(ownersStoreKey)
	store.Set([]byte(name), owner)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(ownersStoreKey)
	bz := store.Get([]byte(name))
	return bz != nil
}
