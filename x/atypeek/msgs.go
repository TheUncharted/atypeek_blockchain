package atypeek

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgAddEndorsement struct {
	IdProject       string
	Contributor     sdk.AccAddress
	ContributorName string
	Receiver        sdk.AccAddress
	ReceiverName    string
	Duration        string
	IdEndorsement   string
	Comments        string
	Skills          string
	Vote            int
}

func (m MsgAddEndorsement) Route() string {
	return "atypeek"
}

func (m MsgAddEndorsement) Type() string {
	return "add_endorsement"
}

func (m MsgAddEndorsement) ValidateBasic() sdk.Error {

	if m.Contributor.Empty() {
		return sdk.ErrInvalidAddress(m.Contributor.String())
	}

	if m.Receiver.Empty() {
		return sdk.ErrInvalidAddress(m.Receiver.String())
	}

	if m.IdProject == "" {
		return sdk.ErrUnknownRequest("IdProject cannot be empty")
	}

	if m.IdEndorsement == "" {
		return sdk.ErrUnknownRequest("IdEndorsement cannot be empty")
	}

	if m.Duration == "" {
		return sdk.ErrUnknownRequest("duration cannot be empty")
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
	return []sdk.AccAddress{m.Contributor}
}

func NewMsgAddEndormsement(idProject string, idEndorsement string, contributor sdk.AccAddress, contributorName string, receiver sdk.AccAddress, receiverName string, duration string, vote int, comments string, skills string) MsgAddEndorsement {
	return MsgAddEndorsement{
		IdProject:       idProject,
		IdEndorsement:   idEndorsement,
		Contributor:     nil,
		ContributorName: contributorName,
		Receiver:        nil,
		ReceiverName:    receiverName,
		Duration:        duration,
		Comments:        comments,
		Skills:          skills,
		Vote:            vote,
	}
}
