package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	ibft "github.com/Mind-chain/mind/command/NLG-ibft"
	"github.com/Mind-chain/mind/command/backup"

	//"github.com/Mind-chain/mind/command/bridge"
	"github.com/Mind-chain/mind/command/genesis"
	"github.com/Mind-chain/mind/command/helper"
	"github.com/Mind-chain/mind/command/license"
	"github.com/Mind-chain/mind/command/monitor"
	"github.com/Mind-chain/mind/command/peers"

	//"github.com/Mind-chain/mind/command/polybft"
	//"github.com/Mind-chain/mind/command/polybftsecrets"
	//"github.com/Mind-chain/mind/command/regenesis"
	//"github.com/Mind-chain/mind/command/rootchain"
	"github.com/Mind-chain/mind/command/secrets"
	"github.com/Mind-chain/mind/command/server"
	"github.com/Mind-chain/mind/command/status"
	"github.com/Mind-chain/mind/command/txpool"
	"github.com/Mind-chain/mind/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "The core node CLI application of Mind Smart Chain",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		//	rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
	//	polybftsecrets.GetCommand(),
	//	polybft.GetCommand(),
	//	bridge.GetCommand(),
	//regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
