package atypeek

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgContrib - high level transaction of the contrib module
type MsgAddProject struct {
	Owner       sdk.AccAddress
	ProjectInfo ProjectInfo
}

func (m MsgAddProject) Route() string {
	return "atypeek"
}

func (m MsgAddProject) Type() string {
	return "add_project"
}

func (m MsgAddProject) ValidateBasic() sdk.Error {

	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.ProjectInfo.Title == "" {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}

	if m.ProjectInfo.StartDate == "" {
		return sdk.ErrUnknownRequest("StartDate cannot be empty")
	}

	if m.ProjectInfo.EndDate == "" {
		return sdk.ErrUnknownRequest("StartDate cannot be empty")
	}

	if m.ProjectInfo.Id == "" {
		return sdk.ErrUnknownRequest("Id cannot be empty")
	}

	return nil
}

func (m MsgAddProject) GetSignBytes() []byte {
	b, err := json.Marshal(m) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	// return b
	return sdk.MustSortJSON(b)
}

func (m MsgAddProject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

func NewMsgAddProject(i ProjectInfo, owner sdk.AccAddress) MsgAddProject {
	return MsgAddProject{
		Owner:       owner,
		ProjectInfo: i,
	}
}
