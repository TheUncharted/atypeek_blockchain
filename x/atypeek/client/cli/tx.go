package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theuncharted/atypeek_blockchain/x/atypeek"
)

const (
	flagId          = "id"
	flagTitle       = "title"
	flagDescription = "description"
	flagStartDate   = "start"
	flagEndDate     = "end"

	flagProjectId     = "projectId"
	flagSkillId       = "skillId"
	flagCourseId      = "courseId"
	flagEndorsementId = "endorsementId"
	flagName          = "name"
	flagTime          = "time"
	flagVote          = "vote"
	flagReceiver      = "receiver"
)

// ContribTxCommand will create a contrib tx and sign it with the given key
func GetCmdAddProject(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-project",
		Short: "Add project to a profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			id := viper.GetString(flagId)
			title := viper.GetString(flagTitle)
			description := viper.GetString(flagDescription)
			startDate := viper.GetString(flagStartDate)
			endDate := viper.GetString(flagEndDate)

			projectInfo := atypeek.ProjectInfo{
				Id:          id,
				Owner:       cliCtx.GetFromAddress(),
				CustomerId:  "",
				Title:       title,
				Description: description,
				StartDate:   startDate,
				EndDate:     endDate,
			}
			msg := atypeek.NewMsgAddProject(projectInfo, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagId, "", "Project Id")
	cmd.Flags().String(flagTitle, "", "Title of project")
	cmd.Flags().String(flagDescription, "", "Description of project")
	cmd.Flags().String(flagStartDate, "", "Start date of project")
	cmd.Flags().String(flagEndDate, "", "End date of project")
	return cmd
}

func GetCmdAddSkill(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-skill",
		Short: "Add skill to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			idProject := viper.GetString(flagProjectId)
			idSkill := viper.GetString(flagSkillId)
			name := viper.GetString(flagName)

			msg := atypeek.NewMsgAddSkill(idProject, idSkill, name, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProjectId, "", "Project Id")
	cmd.Flags().String(flagSkillId, "", "SkillId")
	cmd.Flags().String(flagName, "", "Description of project")
	return cmd
}

func GetCmdAddCourse(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-course",
		Short: "Add course to a skill",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			idCourse := viper.GetString(flagCourseId)
			idSkill := viper.GetString(flagSkillId)
			name := viper.GetString(flagName)

			msg := atypeek.NewMsgAddCourse(idSkill, idCourse, name, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagSkillId, "", "SkillId")
	cmd.Flags().String(flagCourseId, "", "Course Id")
	cmd.Flags().String(flagName, "", "Description of project")
	return cmd
}

func GetCmdAddEndorsement(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-endorsement",
		Short: "Add endorsement to a skill",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			idEndorsement := viper.GetString(flagEndorsementId)
			idSkill := viper.GetString(flagSkillId)
			vote := viper.GetInt(flagVote)
			time := viper.GetString(flagTime)
			dest := viper.GetString(flagReceiver)
			receiver, err := sdk.AccAddressFromBech32(dest)
			if err != nil {
				return err
			}

			msg := atypeek.MsgAddEndorsement{
				Owner:         cliCtx.GetFromAddress(),
				Receiver:      receiver,
				IdSkill:       idSkill,
				IdEndorsement: idEndorsement,
				Time:          time,
				Vote:          vote,
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagSkillId, "", "SkillId")
	cmd.Flags().String(flagEndorsementId, "", "EndorsementId Id")
	cmd.Flags().String(flagReceiver, "", "Receiver adress")
	cmd.Flags().String(flagTime, "", "Date of endorsement")
	cmd.Flags().String(flagVote, "", "Score")
	return cmd
}
