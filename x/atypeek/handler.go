package atypeek

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reflect"
)

// NewHandler returns a handler for "contrib" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddProject:
			return handleAddProject(ctx, k, msg)
		default:
			errMsg := "----Unrecognized Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgContrib.
func handleAddProject(ctx sdk.Context, k Keeper, msg MsgAddProject) sdk.Result {
	resume := k.GetResume(ctx, msg.Owner)
	if resume.Owner.Empty() {
		fmt.Printf("handle add project with no owner \n")
		resume.Owner = msg.Owner
	}
	resume.AddProject(msg.ProjectInfo)
	k.SetResume(ctx, resume)
	return sdk.Result{}
}
