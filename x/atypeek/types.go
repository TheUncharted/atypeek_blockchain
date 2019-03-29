package atypeek

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Course struct {
	Title string         `json:"title"`
	Owner sdk.AccAddress `json:"owner"`
	Xp    int            `json:"xp"`
}

func NewCourse() Course {
	return Course{
		Title: "",
		Owner: nil,
		Xp:    0,
	}
}
