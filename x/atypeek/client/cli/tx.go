package cli

import (
	"fmt"
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
	flagProjectId = "projectId"

	flagEndorsementId = "endorsementId"
	flagName          = "name"
	flagDuration      = "duration"
	flagVote          = "vote"

	flagReceiver     = "receiver"
	flagReceiverName = "receiverName"

	flagContributor     = "contributor"
	flagContributorName = "contributorName"

	flagComments = "comments"
	flagSkills   = "skills"
)

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
			idProject := viper.GetString(flagProjectId)

			c := viper.GetString(flagContributor)
			contributor, err := sdk.AccAddressFromBech32(c)
			if err != nil {
				return err
			}

			contributorName := viper.GetString(flagContributorName)

			r := viper.GetString(flagReceiver)
			receiver, err := sdk.AccAddressFromBech32(r)
			if err != nil {
				return err
			}

			receiverName := viper.GetString(flagReceiverName)

			vote := viper.GetInt(flagVote)
			duration := viper.GetString(flagDuration)

			comments := viper.GetString(flagComments)
			skills := viper.GetString(flagSkills)

			msg := atypeek.MsgAddEndorsement{
				IdProject:       idProject,
				IdEndorsement:   idEndorsement,
				Contributor:     contributor,
				ContributorName: contributorName,
				Receiver:        receiver,
				ReceiverName:    receiverName,
				Duration:        duration,

				Comments: comments,
				Skills:   skills,
				Vote:     vote,
			}

			fmt.Printf("****dfsdfs************ %+v\n", msg)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProjectId, "", "ProjectId")
	cmd.Flags().String(flagEndorsementId, "", "EndorsementId Id")
	cmd.Flags().String(flagContributor, "", "Contributor adress")
	cmd.Flags().String(flagContributorName, "", "Contributor name")
	cmd.Flags().String(flagReceiver, "", "Receiver adress")
	cmd.Flags().String(flagReceiverName, "", "Receiver name")
	cmd.Flags().String(flagDuration, "", "Duration of project")
	cmd.Flags().String(flagVote, "", "Score")
	cmd.Flags().String(flagComments, "", "Comments")
	cmd.Flags().String(flagSkills, "", "Skills")
	return cmd
}
