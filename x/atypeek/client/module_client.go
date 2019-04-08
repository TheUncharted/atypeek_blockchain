package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	nameservicecmd "github.com/theuncharted/atypeek_blockchain/x/atypeek/client/cli"
)

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	namesvcQueryCmd := &cobra.Command{
		Use:   "atypeek",
		Short: "Querying commands for the atypeek module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetCmdResume(mc.storeKey, mc.cdc),
	)...)
	return namesvcQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	namesvcTxCmd := &cobra.Command{
		Use:   "atypeek",
		Short: "Atypeek transactions subcommands",
	}

	namesvcTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.GetCmdAddProject(mc.cdc),
	)...)

	return namesvcTxCmd
}
