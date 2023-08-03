package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:     "restart",
	Short:   "Start or restart one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	rootCmd.AddCommand(restartCmd)

	carapace.Gen(restartCmd).PositionalAnyCompletion(
		action.ActionUnits(restartCmd).FilterArgs(),
	)
}
