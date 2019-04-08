package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper manages transfers between accounts
type Keeper struct {
	cdc      *codec.Codec
	storeKey sdk.StoreKey
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey}
}

func (k Keeper) SetResume(ctx sdk.Context, r Resume) {
	if r.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	jsonstring := string(k.cdc.MustMarshalJSON(r))
	fmt.Printf("set resume with owner %s\n", jsonstring)
	store.Set([]byte(r.Owner.String()), k.cdc.MustMarshalJSON(r))
}

func (k Keeper) GetResume(ctx sdk.Context, owner sdk.AccAddress) Resume {
	fmt.Printf("get resume with owner %s", owner.String())
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(owner.String())) {
		return NewResume()
	}
	bz := store.Get([]byte(owner.String()))
	var resume Resume
	k.cdc.MustUnmarshalJSON(bz, &resume)
	return resume
}
