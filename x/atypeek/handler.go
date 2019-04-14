package atypeek

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reflect"
)

// NewHandler returns a handler for "contrib" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {

		case MsgAddEndorsement:
			return handleAddEndorsement(ctx, k, msg)
		default:
			errMsg := "----Unrecognized Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleAddEndorsement(ctx sdk.Context, k Keeper, msg MsgAddEndorsement) sdk.Result {

	e := Endorsement{
		Id:              msg.IdEndorsement,
		IdProject:       msg.IdProject,
		Contributor:     msg.Contributor,
		ContributorName: msg.ContributorName,
		Receiver:        msg.Receiver,
		ReceiverName:    msg.ReceiverName,
		Duration:        msg.Duration,
		Vote:            msg.Vote,
		Comments:        msg.Comments,
		Skills:          msg.Skills,
	}
	err := k.AddEndorsement(ctx, e, e.Receiver)
	if err != nil {
		errMsg := "add endorsement failed"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.Result{}
}
