package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)

const (
	QueryCourse = "course"
	QueryCouses = "courses"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryCourse:
			return queryCourse(ctx, path[1:], req, keeper)
		case QueryCouses:
			return queryCourses(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown nameservice query endpoint")
		}
	}
}

type QueryResCourse []string

func queryCourses(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	var coursesList QueryResCourse
	iterator := keeper.GetNamesIterator(ctx)

	fmt.Printf("****dsfdsf******* %+v", iterator)
	for ; iterator.Valid(); iterator.Next() {
		course := string(iterator.Key())
		fmt.Printf("***********" + course)
		coursesList = append(coursesList, course)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, coursesList)
	if err2 != nil {
		panic("Could not marshal result to Json")
	}
	return bz, nil

}

func queryCourse(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	title := path[0]
	course := keeper.GetCourse(ctx, title)
	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, course)
	if err2 != nil {
		panic("Could no marshal result to Json")
	}
	return bz, nil
}

func (n QueryResCourse) String() string {
	return strings.Join(n[:], "\n")
}

func (c Course) String() string {
	return strings.TrimSpace(fmt.Sprintf(` Owner: %s Title: %s`, c.Owner, c.Title))
}

func queryBankAccount(ctx sdk.Context, path[]string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return nil, nil
}