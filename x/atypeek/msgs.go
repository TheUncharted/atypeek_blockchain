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

type MsgAddSkill struct {
	Owner     sdk.AccAddress
	IdProject string
	IdSkill   string
	Name      string
}

func (m MsgAddSkill) Route() string {
	return "atypeek"
}

func (m MsgAddSkill) Type() string {
	return "add_skill"
}

func (m MsgAddSkill) ValidateBasic() sdk.Error {

	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.IdProject == "" {
		return sdk.ErrUnknownRequest("IdProject cannot be empty")
	}

	if m.IdSkill == "" {
		return sdk.ErrUnknownRequest("IdSkill cannot be empty")
	}

	if m.Name == "" {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}

	return nil
}

func (m MsgAddSkill) GetSignBytes() []byte {
	b, err := json.Marshal(m) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	// return b
	return sdk.MustSortJSON(b)
}

func (m MsgAddSkill) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

func NewMsgAddSkill(idProject string, idSkill string, name string, owner sdk.AccAddress) MsgAddSkill {
	return MsgAddSkill{
		Owner:     owner,
		IdProject: idProject,
		IdSkill:   idSkill,
		Name:      name,
	}
}

type MsgAddCourse struct {
	Owner    sdk.AccAddress
	IdSkill  string
	IdCourse string
	Name     string
}

func (m MsgAddCourse) Route() string {
	return "atypeek"
}

func (m MsgAddCourse) Type() string {
	return "add_course"
}

func (m MsgAddCourse) ValidateBasic() sdk.Error {

	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.IdCourse == "" {
		return sdk.ErrUnknownRequest("IdCourse cannot be empty")
	}

	if m.IdSkill == "" {
		return sdk.ErrUnknownRequest("IdSkill cannot be empty")
	}

	if m.Name == "" {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}

	return nil
}

func (m MsgAddCourse) GetSignBytes() []byte {
	b, err := json.Marshal(m) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	// return b
	return sdk.MustSortJSON(b)
}

func (m MsgAddCourse) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

func NewMsgAddCourse(idSkill string, idCourse string, name string, owner sdk.AccAddress) MsgAddCourse {
	return MsgAddCourse{
		Owner:    owner,
		IdCourse: idCourse,
		IdSkill:  idSkill,
		Name:     name,
	}
}

type MsgAddEndorsement struct {
	Owner         sdk.AccAddress
	Receiver      sdk.AccAddress
	IdSkill       string
	IdEndorsement string
	Time          string
	Vote          int
}

func (m MsgAddEndorsement) Route() string {
	return "atypeek"
}

func (m MsgAddEndorsement) Type() string {
	return "add_endorsement"
}

func (m MsgAddEndorsement) ValidateBasic() sdk.Error {

	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.Receiver.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.IdSkill == "" {
		return sdk.ErrUnknownRequest("IdSkill cannot be empty")
	}

	if m.IdEndorsement == "" {
		return sdk.ErrUnknownRequest("IdEndorsement cannot be empty")
	}

	if m.Time == "" {
		return sdk.ErrUnknownRequest("time cannot be empty")
	}

	if m.Vote < 0 {
		return sdk.ErrUnknownRequest("vote can not be negative")
	}

	return nil
}

func (m MsgAddEndorsement) GetSignBytes() []byte {
	b, err := json.Marshal(m) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	// return b
	return sdk.MustSortJSON(b)
}

func (m MsgAddEndorsement) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

func NewMsgAddEndormsement(owner sdk.AccAddress, receiver sdk.AccAddress, idSkill string, idEndorsement string, time string, vote int) MsgAddEndorsement {
	return MsgAddEndorsement{
		Owner:         owner,
		Receiver:      receiver,
		IdSkill:       idSkill,
		IdEndorsement: idEndorsement,
		Time:          time,
		Vote:          vote,
	}
}
