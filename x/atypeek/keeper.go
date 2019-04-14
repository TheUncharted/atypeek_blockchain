package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper manages transfers between accounts
type Keeper struct {
	cdc                 *codec.Codec
	profileStoreKey     sdk.StoreKey
	skillStoreKey       sdk.StoreKey
	endorsementStoreKey sdk.StoreKey
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *codec.Codec, profileStoreKey sdk.StoreKey) Keeper {
	return Keeper{cdc: cdc, profileStoreKey: profileStoreKey}
}

func (k Keeper) SetProfile(ctx sdk.Context, r Profile) {
	if r.Owner.Empty() {
		return
	}
	fmt.Printf("set profile with owner %s\n", r.Owner.String())
	store := ctx.KVStore(k.profileStoreKey)
	store.Set([]byte(r.Owner.String()), k.cdc.MustMarshalBinaryBare(r))
}

func (k Keeper) GetProfile(ctx sdk.Context, owner sdk.AccAddress) Profile {
	fmt.Printf("get profile with owner %s", owner.String())
	store := ctx.KVStore(k.profileStoreKey)
	if !store.Has([]byte(owner.String())) {
		profile := NewProfile()
		profile.Owner = owner
		profile.Skills = make(map[string]int)
		return profile
	}
	bz := store.Get([]byte(owner.String()))
	var profile Profile
	k.cdc.MustUnmarshalBinaryBare(bz, &profile)
	return profile
}

func (k Keeper) AddEndorsement(ctx sdk.Context, e Endorsement, receiver sdk.AccAddress) error {
	profile := k.GetProfile(ctx, receiver)
	profile.Endorsements = append(profile.Endorsements, e)

	//var skills []string
	//if len(e.Skills) > 0 {
	//	skills = strings.Split(e.Skills, ",")
	//	for _, skill := range skills {
	//		skill = strings.TrimSpace(skill)
	//		fmt.Println("skill %s", skill)
	//
	//		if val, ok := profile.Skills[skill]; ok {
	//			fmt.Println("+++++")
	//			profile.Skills[skill] = val + e.Vote
	//		} else {
	//			fmt.Println("-----")
	//			profile.Skills[skill] = e.Vote
	//		}
	//
	//	}
	//
	//}

	k.SetProfile(ctx, profile)
	return nil
}
