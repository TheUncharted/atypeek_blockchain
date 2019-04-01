package atypeek

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type Contrib interface {
	GetKey() []byte
	GetContributor() sdk.AccAddress
	GetTime() string
	AppendTags(*sdk.Tags)
	NewStatus() Status
	ValidateBasic() sdk.Error
	ValidateAccounts(sdk.Context, auth.AccountKeeper) (auth.Account, sdk.Error)
	String() string
}

type Contribs []Contrib

// ValidateBasic - validate transaction contribs
func (contribs Contribs) ValidateBasic() sdk.Error {
	// m := make(map[string]struct{})
	for _, ctb := range contribs {
		err := ctb.ValidateBasic()
		if err != nil {
			return err
		}
		// _, found := m[string(ctb.Key)]
		// if found {
		// 	return ErrInvalidContrib(DefaultCodespace, "duplicate key")
		// }
		// m[string(ctb.Key)] = struct{}{}
	}

	return nil
}

type BaseContrib struct {
	Key         []byte         `json:"key"`
	Contributor sdk.AccAddress `json:"contributor"`
	Time        string         `json:"time"`
}

// Implements Contrib
func (ctb BaseContrib) GetKey() []byte {
	return ctb.Key
}

func (ctb BaseContrib) GetContributor() sdk.AccAddress {
	return ctb.Contributor
}

func (ctb BaseContrib) GetTime() string {
	return ctb.Time
}

func (ctb BaseContrib) AppendTags(tags *sdk.Tags) {
	*tags = append(*tags, sdk.MakeTag("contributor", ctb.Contributor.String()))
}

func (ctb BaseContrib) NewStatus() Status {
	return &BaseStatus{Score: 1, Contributor: ctb.Contributor, Time: ctb.Time}
}

func (ctb BaseContrib) ValidateBasic() sdk.Error {
	if len(ctb.Key) == 0 {
		return ErrInvalidContrib(DefaultCodespace, ctb.String())
	}
	if len(ctb.Contributor) == 0 {
		return sdk.ErrInvalidAddress(ctb.Contributor.String())
	}
	return nil
}

func (ctb BaseContrib) ValidateAccounts(ctx sdk.Context, am auth.AccountKeeper) (auth.Account, sdk.Error) {
	acc := am.GetAccount(ctx, ctb.Contributor)
	if acc == nil {
		fmt.Println("Unknown address")
		return nil, sdk.ErrUnknownAddress(ctb.Contributor.String())
	}
	return acc, nil
}

func (ctb BaseContrib) String() string {
	return fmt.Sprintf("%v", ctb)
}

type BaseContrib2 struct {
	BaseContrib
	Recipient sdk.AccAddress `json:"recipient"`
}

func (ctb BaseContrib2) AppendTags(tags *sdk.Tags) {
	*tags = append(*tags, sdk.MakeTag("contributor", ctb.Contributor.String()), sdk.MakeTag("recipient", ctb.Recipient.String()))
}

func (ctb BaseContrib2) NewStatus() Status {
	return &BaseStatus2{BaseStatus: BaseStatus{Score: 1, Contributor: ctb.Contributor, Time: ctb.Time}, Recipient: ctb.Recipient}
}

func (ctb BaseContrib2) ValidateBasic() sdk.Error {
	if len(ctb.Key) == 0 {
		return ErrInvalidContrib(DefaultCodespace, ctb.String())
	}
	if len(ctb.Contributor) == 0 {
		return sdk.ErrInvalidAddress(ctb.Contributor.String())
	}
	if len(ctb.Recipient) == 0 {
		return sdk.ErrInvalidAddress(ctb.Recipient.String())
	}
	return nil
}

func (ctb BaseContrib2) ValidateAccounts(ctx sdk.Context, am auth.AccountKeeper) (auth.Account, sdk.Error) {
	acc := am.GetAccount(ctx, ctb.Contributor)
	if acc == nil {
		return nil, sdk.ErrUnknownAddress(ctb.Contributor.String())
	}

	if am.GetAccount(ctx, ctb.Recipient) == nil {
		return nil, sdk.ErrUnknownAddress(ctb.Recipient.String())
	}

	return acc, nil
}

// Status - contrib status
type Status interface {
	GetScore() int64
	Update(Contrib) sdk.Error
}

type BaseStatus struct {
	Score       int64          `json:"score"`
	Contributor sdk.AccAddress `json:"contributor"`
	Time        string         `json:"time"`
}

func (status BaseStatus) GetScore() int64 {
	return status.Score
}

func (status *BaseStatus) Update(ctb Contrib) sdk.Error {
	// check if addr is the contributor
	if !bytes.Equal(ctb.GetContributor(), status.Contributor) {
		return sdk.ErrUnknownAddress("contributor error")
	}
	// check if time is valid
	//if !ctb.GetTime().After(status.Time) {
	//	return sdk.ErrUnknownAddress("time error")
	//}
	// TODO: better score calculation
	status.Score++
	status.Time = ctb.GetTime()
	return nil
}

type BaseStatus2 struct {
	BaseStatus
	Recipient sdk.AccAddress `json:"recipient"`
}

func (status *BaseStatus2) Update(ctb Contrib) sdk.Error {
	ctb2 := ctb.(BaseContrib2)
	// check if addr is the contributor
	if !bytes.Equal(ctb2.Contributor, status.Contributor) {
		return sdk.ErrUnknownAddress("contributor error")
	}
	// check if the recipient is matched
	if !bytes.Equal(ctb2.Recipient, status.Recipient) {
		return sdk.ErrUnknownAddress("recipient error")
	}
	// check if time is valid
	//if !ctb2.Time.After(status.Time) {
	//	return sdk.ErrUnknownAddress("time error")
	//}
	// TODO: better score calculation
	status.Score++
	status.Time = ctb.GetTime()
	return nil
}

type Invite struct {
	BaseContrib2
	Content []byte `json:"content"`
}

func (ctb Invite) NewStatus() Status {
	return &InviteStatus{BaseStatus: BaseStatus{Score: 1, Contributor: ctb.Contributor, Time: ctb.Time}, Recipient: ctb.Recipient}
}

func (ctb Invite) ValidateAccounts(ctx sdk.Context, am auth.AccountKeeper) (auth.Account, sdk.Error) {
	fmt.Println("conributor Invite %v", ctb.Contributor.String())
	fmt.Println("Recipient Invite %v", ctb.Recipient.String())

	acc := am.GetAccount(ctx, ctb.Contributor)
	fmt.Println("account Invite %v", acc.String())
	if acc == nil {
		fmt.Println("no accout contrib")
		return nil, sdk.ErrUnknownAddress(ctb.Contributor.String())
	}
	if am.GetAccount(ctx, ctb.Recipient) != nil {
		fmt.Println("no accout Recipient")
		return nil, sdk.ErrUnknownAddress(ctb.Recipient.String())
	}

	//am.SetAccount(ctx, am.NewAccountWithAddress(ctx, ctb.Recipient))

	return acc, nil
}

type InviteStatus BaseStatus2

func (status *InviteStatus) Update(ctb Contrib) sdk.Error {
	ctb2 := ctb.(*Invite)
	// check if addr is the contributor
	if !bytes.Equal(ctb2.Contributor, status.Contributor) {
		return sdk.ErrUnknownAddress("contributor error")
	}
	// check if the recipient is matched
	if !bytes.Equal(ctb2.Recipient, status.Recipient) {
		return sdk.ErrUnknownAddress("recipient error")
	}
	// check if time is valid
	//if !ctb2.Time.After(status.Time) {
	//	return sdk.ErrUnknownAddress("time error")
	//}

	status.Time = ctb.GetTime()
	return nil
}

type Recommend struct {
	BaseContrib2
	Content []byte `json:"content"`
}

type RecommendStatus BaseStatus2

func (ctb Recommend) NewStatus() Status {
	return &RecommendStatus{BaseStatus: BaseStatus{Score: 1, Contributor: ctb.Contributor, Time: ctb.Time}, Recipient: ctb.Recipient}
}

func (ctb Recommend) ValidateAccounts(ctx sdk.Context, am auth.AccountKeeper) (auth.Account, sdk.Error) {
	acc := am.GetAccount(ctx, ctb.Contributor)
	if acc == nil {
		return nil, sdk.ErrUnknownAddress(ctb.Contributor.String())
	}
	if am.GetAccount(ctx, ctb.Recipient) == nil {
		return nil, sdk.ErrUnknownAddress(ctb.Recipient.String())
	}
	// cannot recommend his/herself
	if acc == am.GetAccount(ctx, ctb.Recipient) {
		return nil, sdk.ErrInvalidAddress(ctb.Recipient.String())
	}

	return acc, nil
}

type Post struct {
	BaseContrib2
	Content []byte `json:"content"`
}

type PostStatus BaseStatus2

func (ctb Post) NewStatus() Status {
	return &PostStatus{BaseStatus: BaseStatus{Score: 1, Contributor: ctb.Contributor, Time: ctb.Time}, Recipient: ctb.Recipient}
}
