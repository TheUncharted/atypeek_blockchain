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

func GetCmdProject(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "project [id]",
		Short: "Query project ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/project/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not resolve project - %s \n", id)
				return nil
			}

			var out atypeek.Project
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdSkill(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "skill [id]",
		Short: "Query skill ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/skill/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not resolve skill - %s \n", id)
				return nil
			}

			var out atypeek.Skill
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCourse(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "course [id]",
		Short: "Query course ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/course/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not resolve course - %s \n", id)
				return nil
			}

			var out atypeek.Course
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
