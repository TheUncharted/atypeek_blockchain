package atypeek

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*IProfile)(nil), nil)
	cdc.RegisterConcrete(Profile{}, "atypeek/Profile", nil)
	cdc.RegisterConcrete(Project{}, "atypeek/Project", nil)
	cdc.RegisterConcrete(ProjectInfo{}, "atypeek/ProjectInfo", nil)
	cdc.RegisterConcrete(MsgAddProject{}, "nameservice/AddProject", nil)
	cdc.RegisterConcrete(Skill{}, "atypeek/Skill", nil)
	cdc.RegisterConcrete(MsgAddSkill{}, "nameservice/AddSkill", nil)

	cdc.RegisterConcrete(Course{}, "atypeek/Course", nil)
	cdc.RegisterConcrete(MsgAddCourse{}, "nameservice/AddCourse", nil)

	cdc.RegisterConcrete(Endorsement{}, "atypeek/Endorsement", nil)
	cdc.RegisterConcrete(MsgAddEndorsement{}, "nameservice/AddEndorsement", nil)

}
