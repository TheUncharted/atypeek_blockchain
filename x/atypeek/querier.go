package atypeek

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryProfile     = "profile"
	QueryProject     = "project"
	QuerySkill       = "skill"
	QueryCourse      = "course"
	QueryEndorsement = "endorsement"
	QuerySkillScore  = "score"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryProfile:
			return queryProfile(ctx, path[1:], req, keeper)
		case QueryProject:
			return queryProject(ctx, path[1:], req, keeper)
		case QuerySkill:
			return querySkill(ctx, path[1:], req, keeper)
		case QueryCourse:
			return queryCourse(ctx, path[1:], req, keeper)
		case QueryEndorsement:
			return queryEndorsement(ctx, path[1:], req, keeper)
		case QuerySkillScore:
			return querySkillScore(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryProfile(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	address := path[0]
	owner, err2 := sdk.AccAddressFromBech32(address)
	if err2 != nil {
		panic("could not find owner")
	}

	profile := keeper.GetProfile(ctx, owner)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, profile)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryProject(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	project, err2 := keeper.GetProject(ctx, id)
	if err2 == nil {
		bz, err2 := codec.MarshalJSONIndent(keeper.cdc, project)
		if err2 != nil {
			panic("could not marshal result to JSON")
		}

		return bz, nil
	}
	panic("could not find project")
}

func querySkill(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	skill, err2 := keeper.GetSkill(ctx, id)
	if err2 == nil {
		bz, err2 := codec.MarshalJSONIndent(keeper.cdc, skill)
		if err2 != nil {
			panic("could not marshal result to JSON")
		}

		return bz, nil
	}
	panic("could not find skill")
}

func queryCourse(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	course, err2 := keeper.GetCourse(ctx, id)
	if err2 == nil {
		bz, err2 := codec.MarshalJSONIndent(keeper.cdc, course)
		if err2 != nil {
			panic("could not marshal result to JSON")
		}

		return bz, nil
	}
	panic("could not find skill")
}

func queryEndorsement(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	fmt.Printf("*********skill queryEndorsement", id)
	endorsement, err2 := keeper.GetEndorsement(ctx, id)
	if err2 == nil {
		bz, err2 := codec.MarshalJSONIndent(keeper.cdc, endorsement)
		if err2 != nil {
			panic("could not marshal result to JSON")
		}

		return bz, nil
	}
	panic("could not find endorsement")
}

func querySkillScore(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	fmt.Printf("*********skill score", id)
	skillScore := SkillScore{
		IdSkill: id,
		Score:   0,
	}

	skill, err2 := keeper.GetSkill(ctx, id)
	if err2 == nil && len(skill.Endorsements) > 0 {
		for _, idEndorsement := range skill.Endorsements {
			endorsement, err2 := keeper.GetEndorsement(ctx, idEndorsement)
			if err2 == nil {
				skillScore.Score += endorsement.Vote
			}
		}

	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, skillScore)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil

}
