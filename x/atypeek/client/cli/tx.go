package cli

import (
	"encoding/hex"
	"errors"
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
	flagTo      = "to"
	flagKey     = "key"
	flagType    = "type"
	flagContent = "content"
	flagVotes   = "votes"
	flagTime    = "time"
	// flagRole = "role"
	// flagAsync  = "async"
)

func BuildContribMsg(ctb atypeek.Contrib) sdk.Msg {
	msg := atypeek.NewMsgContrib(atypeek.Contribs{ctb})
	return msg
}

// ContribTxCommand will create a contrib tx and sign it with the given key
func ContribTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contrib",
		Short: "Create and sign a contrib tx",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// get the from address
			from := cliCtx.GetFromAddress()

			ctbKey, err := hex.DecodeString(viper.GetString(flagKey))
			if err != nil {
				return err
			}

			ctbTime := viper.GetString(flagTime)

			ctbContent := []byte(viper.GetString(flagContent))
			fmt.Printf("content flag = %s", string(ctbContent))

			// parse destination address
			dest := viper.GetString(flagTo)
			to, err := sdk.AccAddressFromBech32(dest)
			if err != nil {
				return err
			}

			ctbType := viper.GetString(flagType)
			var ctb atypeek.Contrib
			switch ctbType {
			case "Invite", "Recommend", "Post":

				switch ctbType {
				case "Invite":
					ctb = atypeek.Invite{atypeek.BaseContrib2{atypeek.BaseContrib{ctbKey, from, ctbTime}, to}, ctbContent}
				case "Post":
					ctb = atypeek.Post{atypeek.BaseContrib2{atypeek.BaseContrib{ctbKey, from, ctbTime}, to}, ctbContent}
				case "Recommend":
					ctb = atypeek.Recommend{atypeek.BaseContrib2{atypeek.BaseContrib{ctbKey, from, ctbTime}, to}, ctbContent}
				}

			default:
				return errors.New("Invalid Contrib Type")
			}

			msg := BuildContribMsg(ctb)
			cliCtx.PrintResponse = true
			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTo, "", "Address to contrib")
	cmd.Flags().String(flagKey, "", "Key of the contrib")
	cmd.Flags().String(flagType, "", "Type of the contrib")
	cmd.Flags().String(flagContent, "", "Content of the contrib")
	cmd.Flags().String(flagVotes, "", "Votes of the contrib")
	cmd.Flags().String(flagTime, "", "Time of the contrib")
	// cmd.Flags().Bool(flagAsync, false, "Pass the async flag to send a tx without waiting for the tx to be included in a block")
	return cmd
}
