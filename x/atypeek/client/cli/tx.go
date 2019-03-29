package cli

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/theuncharted/cosmostest/x/atypeek"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdBuyName is the CLI command for sending a BuyName transaction
func GetCmdAddCourse(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-course title",
		Short: "add course in blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			title := args[0]
			if title == "" {
				return errors.New("title is empty")
			}

			msg := atypeek.NewMsgAddCourse(title, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

func GetCmdBankAccountEvent(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-account-event",
		Short: "add event in bank account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			event := args[0]
			if event == "" {
				return errors.New("event is empty")
			}

			amount, err := strconv.ParseInt(args[1], 0, 64)
			if err != nil || amount <= 0 {
				return errors.New("amount must be positive")
			}

			msg := atypeek.NewMsgBankAccountEvent(event, amount, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}
