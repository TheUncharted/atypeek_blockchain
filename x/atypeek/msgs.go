package atypeek

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgAddCourse struct {
	Title string
	Owner sdk.AccAddress
}

func NewMsgAddCourse(title string, owner sdk.AccAddress) MsgAddCourse {
	return MsgAddCourse{
		Title: title,
		Owner: owner,
	}
}

func (msg MsgAddCourse) Route() string { return "nameservice" }
func (msg MsgAddCourse) Type() string  { return "add_course" }

func (msg MsgAddCourse) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Title) == 0 {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}

	return nil
}

func (msg MsgAddCourse) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgAddCourse) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgBankAccountEvent struct {
	Event  string
	Amount int64
	Owner  sdk.AccAddress
}

func NewMsgBankAccountEvent(event string, amount int64, owner sdk.AccAddress) MsgBankAccountEvent {
	return MsgBankAccountEvent{
		Event:  event,
		Amount: amount,
		Owner:  owner,
	}
}

func (msg MsgBankAccountEvent) Route() string {
	return "atypeek"
}

func (msg MsgBankAccountEvent) Type() string {
	return "event"
}

func (msg MsgBankAccountEvent) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if msg.Event != "deposit" && msg.Event != "withdraw" {
		return sdk.ErrUnknownRequest("Event not authorized")
	}

	if msg.Amount <= 0 {
		return sdk.ErrUnknownRequest("Amount must be positive")
	}

	return nil
}

func (msg MsgBankAccountEvent) GetSignBytes() []byte {

	b, err := json.Marshal(msg)

	if err != nil {

		panic(err)
	}

	return sdk.MustSortJSON(b)
}

func (msg MsgBankAccountEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
