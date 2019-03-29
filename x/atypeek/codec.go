package atypeek

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddCourse{}, "nameservice/AddCourse", nil)
	cdc.RegisterConcrete(MsgBankAccountEvent{}, "nameservice/BankAccountEvent", nil)

}