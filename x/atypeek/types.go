package atypeek

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Profile struct {
		Owner        sdk.AccAddress `json:"owner"`
		Endorsements []Endorsement  `json:"endorsements"`
		Skills       map[string]int `json:"skills"`
	}

	Endorsement struct {
		Id              string         `json:"id"`
		IdProject       string         `json:"idProject"`
		Contributor     sdk.AccAddress `json:"contributor"`
		ContributorName string         `json:"contributorName"`
		Receiver        sdk.AccAddress `json:"receiver"`
		ReceiverName    string         `json:"receiverName"`
		Duration        string         `json:"duration"`
		Vote            int            `json:"vote"`
		Comments        string         `json:"comments"`
		Skills          string         `json:"skills"`
	}
)

func NewProfile() Profile {
	return Profile{
		Owner:        nil,
		Endorsements: []Endorsement{},
		Skills:       make(map[string]int),
	}
}

func (p Profile) String() string {
	return fmt.Sprintf("%+v", p)
}

func (e Endorsement) String() string {
	return fmt.Sprintf("%+v", e)
}

func NewEndorsement() Endorsement {
	return Endorsement{
		Id:              "",
		IdProject:       "",
		Contributor:     nil,
		ContributorName: "",
		Receiver:        nil,
		ReceiverName:    "",
		Duration:        "",
		Vote:            0,
		Comments:        "",
		Skills:          "",
	}
}

type SkillScore struct {
	IdSkill string `json:"id"`
	Score   int    `json:"score"`
}

func (s SkillScore) String() string {
	return fmt.Sprintf("%+v", s)
}
