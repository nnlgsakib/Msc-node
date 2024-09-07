package ibft

import (
	"github.com/Mind-chain/mind/command/NLG-ibft/candidates"
	"github.com/Mind-chain/mind/command/NLG-ibft/propose"
	"github.com/Mind-chain/mind/command/NLG-ibft/quorum"
	"github.com/Mind-chain/mind/command/NLG-ibft/snapshot"
	"github.com/Mind-chain/mind/command/NLG-ibft/status"
	_switch "github.com/Mind-chain/mind/command/NLG-ibft/switch"
	"github.com/Mind-chain/mind/command/helper"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "nibft",
		Short: "Top level NLG-IBFT command for interacting with the NLG-IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
