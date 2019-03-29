package atypeek

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddCourse:
			return handleMsgAddCourse(ctx, keeper, msg)
		case MsgBankAccountEvent:
			return handleMsgBankAccountEvent(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg Type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgBankAccountEvent(ctx sdk.Context, keeper Keeper, msg MsgBankAccountEvent) sdk.Result {
	if msg.Owner.Empty() {
		return sdk.ErrUnknownRequest("Owner must not be empty").Result()
	}

	if msg.Amount <= 0 {
		return sdk.ErrUnknownRequest("Amount must be positive").Result()
	}

	coin := sdk.NewCoin("atk", sdk.NewInt(msg.Amount))

	switch msg.Event {
	case "deposit":
		keeper.Deposit(ctx, msg.Owner, sdk.Coins{coin})
	case "withdraw":
		keeper.Withdraw(ctx, msg.Owner, sdk.Coins{coin})
	default:
		errMsg := fmt.Sprintf("Unrecognized event Msg Event %v", msg.Event)
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
	return sdk.Result{}
}

func handleMsgAddCourse(ctx sdk.Context, keeper Keeper, msg MsgAddCourse) sdk.Result {
	if msg.Owner.Empty() {
		return sdk.ErrUnknownRequest("Owner must not be empty").Result()
	}
	keeper.SetOwner(ctx, msg.Title, msg.Owner)
	keeper.AddCourse(ctx, msg.Title)
	return sdk.Result{}
}
