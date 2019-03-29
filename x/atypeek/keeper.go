package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)



type Keeper struct {
	courseStoreKey      sdk.StoreKey
	bankKeeper  bank.Keeper
	cdc *codec.Codec
}

func NewKeeper(courseStoreKey sdk.StoreKey,   cdc *codec.Codec) Keeper {
	return Keeper{
		courseStoreKey:      courseStoreKey,
		cdc:                 cdc,
	}
}




func (k Keeper) AddCourse(ctx sdk.Context, title string) {
	course := k.GetCourse(ctx, title)
	course.Title = title
	k.SetCourse(ctx, title, course)

}

func (k Keeper) SetCourse(ctx sdk.Context, title string, course Course) {
	if course.Owner.Empty() {
		fmt.Printf("No owner")
		return
	}
	store := ctx.KVStore(k.courseStoreKey)
	store.Set([]byte(title), k.cdc.MustMarshalBinaryBare(course))

}

func (k Keeper) GetCourse(ctx sdk.Context, title string) Course {
	store := ctx.KVStore(k.courseStoreKey)
	if !store.Has([]byte(title)) {
		return NewCourse()
	}
	bz := store.Get([]byte(title))
	var course Course
	k.cdc.MustUnmarshalBinaryBare(bz, &course)
	return course
}

func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.courseStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, title string) bool {
	return !k.GetCourse(ctx, title).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, title string) sdk.AccAddress {
	return k.GetCourse(ctx, title).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, title string, owner sdk.AccAddress) {
	course := k.GetCourse(ctx, title)
	course.Owner = owner
	k.SetCourse(ctx, title, course)
}

func (k Keeper) Deposit(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins) {

	k.bankKeeper.AddCoins(ctx, owner, coins)
}

func (k Keeper) Withdraw(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins ) {
	k.bankKeeper.SubtractCoins(ctx, owner, coins)
}