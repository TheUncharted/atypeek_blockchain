package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// Keeper manages transfers between accounts
type Keeper struct {
	cdc      *codec.Codec
	am       auth.AccountKeeper
	storeKey sdk.StoreKey
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *codec.Codec, am auth.AccountKeeper, storeKey sdk.StoreKey) Keeper {
	return Keeper{cdc: cdc, am: am, storeKey: storeKey}
}

func (k Keeper) UpdateContrib(ctx sdk.Context, ctb Contrib, tags *sdk.Tags) sdk.Error {

	fmt.Println("ValidateAccounts")
	acc, err := ctb.ValidateAccounts(ctx, k.am)
	if err != nil {
		fmt.Println("----------------------------------")
		return err
	}

	fmt.Println("append tags")
	ctb.AppendTags(tags)

	var oldscore int64
	fmt.Println("store  key %v", k.storeKey)
	store := ctx.KVStore(k.storeKey)
	key := ctb.GetKey()

	fmt.Println("getStatus")
	status, err := getStatus(store, key, k.cdc)
	if err != nil {
		return err
	}
	if status != nil {
		oldscore = status.GetScore()
		err := status.Update(ctb)
		if err != nil {
			return err
		}
	} else {
		oldscore = 0
		status = ctb.NewStatus()
	}
	fmt.Println("GetScore")
	diff := status.GetScore() - oldscore
	setStatus(store, key, status, k.cdc)
	fmt.Println("update reput")
	updateRepute(acc, diff)
	fmt.Println("update set account")
	k.am.SetAccount(ctx, acc)

	return nil
}

func getStatus(store sdk.KVStore, key []byte, cdc *codec.Codec) (Status, sdk.Error) {
	var status Status
	data := store.Get(key)
	if len(data) > 0 {
		err := cdc.UnmarshalBinaryBare(data, &status)
		if err != nil {
			// msg := fmt.Sprintf("Error reading contrib %X", key)
			return nil, sdk.ErrUnknownAddress(fmt.Sprintf("%s", err))
		}
		return status, nil
	}
	return nil, nil
}

func setStatus(store sdk.KVStore, key []byte, status Status, cdc *codec.Codec) {
	bin, _ := cdc.MarshalBinaryBare(status)
	store.Set(key, bin)
}

func updateRepute(acc auth.Account, diff int64) {
	//acc.(*types.ReputeAccount).Repute += diff
}
