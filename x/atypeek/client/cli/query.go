package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/cobra"
	"github.com/theuncharted/atypeek_blockchain/x/atypeek"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdResume(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resume [user account]",
		Short: "Query resume info of user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			ownerAddress := args[0]

			fmt.Printf("query route %+v", queryRoute)
			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resume/%s", queryRoute, ownerAddress), nil)
			if err != nil {
				fmt.Printf("could not resolve resume - %s \n", ownerAddress)
				return nil
			}

			var out atypeek.Resume
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
