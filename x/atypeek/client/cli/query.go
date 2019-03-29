package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/theuncharted/cosmostest/x/atypeek"
)

func GetCmdCourse(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "course [title]",
		Short: "Query course info of title",
		Args: cobra.ExactArgs(1),
		RunE:func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			title := args[0]
			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/course/%s", queryRoute, title), nil)
			if err != nil {
				fmt.Printf("could not resolve course - %s \n", string(title))
			}
			var out atypeek.Course
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdCourses(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "courses",
		Short: "Courses",
		//Args: cobra.ExactArgs(1),
		RunE:func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/courses", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not query courses - %s \n")
			}
			var out atypeek.QueryResCourse
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
