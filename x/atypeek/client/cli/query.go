package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/cobra"
	"github.com/theuncharted/atypeek_blockchain/x/atypeek"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdProfile(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "profile [user account]",
		Short: "Query profile info of user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			ownerAddress := args[0]

			fmt.Printf("query route %+v", queryRoute)
			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/profile/%s", queryRoute, ownerAddress), nil)
			if err != nil {
				fmt.Printf("could not resolve profile - %s \n", ownerAddress)
				return nil
			}

			var out atypeek.Profile
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdEndorsement(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "endorsement [id]",
		Short: "Query endorsement ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/endorsement/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not resolve endorsement - %s \n", id)
				return nil
			}

			var out atypeek.Endorsement
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdSkillScore(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "score [id]",
		Short: "Query skill score ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/score/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not resolve skill score - %s \n", id)
				return nil
			}

			var out atypeek.SkillScore
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
