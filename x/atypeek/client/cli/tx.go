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
)

// ContribTxCommand will create a contrib tx and sign it with the given key
func GetCmdAddProject(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-project",
		Short: "Add project to a resume",
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
