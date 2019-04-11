package atypeek

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reflect"
)

// NewHandler returns a handler for "contrib" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddProject:
			return handleAddProject(ctx, k, msg)
		case MsgAddSkill:
			return handleAddSkill(ctx, k, msg)
		case MsgAddCourse:
			return handleAddCourse(ctx, k, msg)
		case MsgAddEndorsement:
			return handleAddEndorsement(ctx, k, msg)
		default:
			errMsg := "----Unrecognized Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgContrib.
func handleAddProject(ctx sdk.Context, k Keeper, msg MsgAddProject) sdk.Result {
	//profile := k.GetProfile(ctx, msg.Owner)
	//if profile.Owner.Empty() {
	//	fmt.Printf("handle add project with no owner \n")
	//	profile.Owner = msg.Owner
	//}

	project := NewProjectWithProjectInfo(msg.ProjectInfo)
	err := k.AddProject(ctx, project, msg.Owner)
	if err != nil {
		errMsg := "add project failed"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.Result{}
}

func handleAddSkill(ctx sdk.Context, k Keeper, msg MsgAddSkill) sdk.Result {

	skill := Skill{
		Owner:   msg.Owner,
		Id:      msg.IdSkill,
		Name:    msg.Name,
		Courses: nil,
	}
	err := k.AddSkill(ctx, msg.IdProject, skill, msg.Owner)
	if err != nil {
		errMsg := "add skill failed"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.Result{}
}

func handleAddCourse(ctx sdk.Context, k Keeper, msg MsgAddCourse) sdk.Result {

	course := Course{
		Owner: msg.Owner,
		Id:    msg.IdCourse,
		Name:  msg.Name,
	}
	err := k.AddCourse(ctx, msg.IdSkill, course, msg.Owner)
	if err != nil {
		errMsg := "add skill failed"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.Result{}
}

func handleAddEndorsement(ctx sdk.Context, k Keeper, msg MsgAddEndorsement) sdk.Result {

	endorsement := Endorsement{
		Id:          msg.IdEndorsement,
		IdSkill:     msg.IdSkill,
		Contributor: msg.Owner,
		Time:        msg.Time,
		Receiver:    msg.Receiver,
		Vote:        msg.Vote,
	}
	err := k.AddEndorsement(ctx, msg.IdSkill, endorsement, msg.Owner)
	if err != nil {
		errMsg := "add endorsement failed"
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.Result{}
}
