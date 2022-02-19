package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/mt/exported"
	"github.com/irisnet/irismod/modules/mt/types"
)

// GetMT gets the the specified MT
func (k Keeper) GetMT(ctx sdk.Context, denomID, tokenID string) (mt exported.MT, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyMT(denomID, tokenID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found MT: %s", denomID)
	}

	var baseMT types.MT
	k.cdc.MustUnmarshal(bz, &baseMT)

	return baseMT, nil
}

// GetMTs returns all MTs by the specified denom ID
func (k Keeper) GetMTs(ctx sdk.Context, denom string) (mts []exported.MT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyMT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseMT types.MT
		k.cdc.MustUnmarshal(iterator.Value(), &baseMT)
		mts = append(mts, baseMT)
	}

	return mts
}

// Authorize checks if the sender is the owner of the given MT
// Return the MT if true, an error otherwise
func (k Keeper) Authorize(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) (types.MT, error) {
	mt, err := k.GetMT(ctx, denomID, tokenID)
	if err != nil {
		return types.MT{}, err
	}

	if !owner.Equals(mt.GetOwner()) {
		return types.MT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}

	return mt.(types.MT), nil
}

// CheckMt heck whether MT can be transferred
func (k Keeper) CheckMt(ctx sdk.Context, denomID, tokenID string, amount uint64, owner sdk.AccAddress) (bool, error) {
	if k.HasMT(ctx, denomID, tokenID) {
		return false, sdkerrors.Wrapf(types.ErrUnknownMT, "MT %s unknown in denom %s", tokenID, denomID)
	}

	balanceAmount := k.getBalance(ctx, owner, denomID, tokenID).Amount
	if amount-balanceAmount < 0 {
		return false, sdkerrors.Wrapf(types.ErrInvalidTokenAmount, "%s MT balance amount %v", owner.String(), balanceAmount)
	}

	return true, nil
}

// HasMT checks if the specified MT exists
func (k Keeper) HasMT(ctx sdk.Context, denomID, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyMT(denomID, tokenID))
}

func (k Keeper) setMT(ctx sdk.Context, denomID string, mt types.MT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&mt)
	store.Set(types.KeyMT(denomID, mt.GetID()), bz)
}

// deleteMT deletes an existing MT from store
func (k Keeper) deleteMT(ctx sdk.Context, denomID string, mt exported.MT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyMT(denomID, mt.GetID()))
}
