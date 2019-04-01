package atypeek

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Contrib)(nil), nil)
	cdc.RegisterConcrete(&Invite{}, "atypeek/Invite", nil)
	cdc.RegisterConcrete(&Recommend{}, "atypeek/Recommend", nil)
	//cdc.RegisterConcrete(&Vote{}, "contrib/Vote", nil)
	cdc.RegisterConcrete(&Post{}, "contrib/Post", nil)
	cdc.RegisterInterface((*Status)(nil), nil)
	cdc.RegisterConcrete(&InviteStatus{}, "atypeek/InviteStatus", nil)
	cdc.RegisterConcrete(&RecommendStatus{}, "atypeek/RecommendStatus", nil)
	//cdc.RegisterConcrete(&VoteStatus{}, "contrib/VoteStatus", nil)
	cdc.RegisterConcrete(&PostStatus{}, "contrib/PostStatus", nil)
	cdc.RegisterConcrete(MsgContrib{}, "forbole/ContribMsg", nil)

}
