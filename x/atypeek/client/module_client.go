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
		Use:   "nameservice",
		Short: "Querying commands for the nameservice module",
	}

	namesvcQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetContribCmd(mc.storeKey, mc.cdc),
	)...)
	return namesvcQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	namesvcTxCmd := &cobra.Command{
		Use:   "atypeek",
		Short: "Nameservice transactions subcommands",
	}

	namesvcTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.ContribTxCmd(mc.cdc),
	)...)

	return namesvcTxCmd
}
