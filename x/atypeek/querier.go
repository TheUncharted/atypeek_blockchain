package atypeek

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryResume = "resume"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResume:
			return queryResume(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryResume(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	address := path[0]
	owner, err2 := sdk.AccAddressFromBech32(address)
	if err2 != nil {
		panic("could not find owner")
	}

	resume := keeper.GetResume(ctx, owner)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, resume)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
