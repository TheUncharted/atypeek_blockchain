package atypeek

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*IResume)(nil), nil)
	cdc.RegisterConcrete(Resume{}, "atypeek/Resume", nil)
	cdc.RegisterConcrete(Project{}, "atypeek/Project", nil)
	cdc.RegisterConcrete(ProjectInfo{}, "atypeek/ProjectInfo", nil)
	cdc.RegisterConcrete(MsgAddProject{}, "nameservice/AddProject", nil)

}
