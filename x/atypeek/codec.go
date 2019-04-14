package atypeek

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(Profile{}, "atypeek/Profile", nil)

	cdc.RegisterConcrete(Endorsement{}, "atypeek/Endorsement", nil)
	cdc.RegisterConcrete(MsgAddEndorsement{}, "nameservice/AddEndorsement", nil)

}
