package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resetFailedCmd = &cobra.Command{
	Use:     "reset-failed",
	Short:   "Reset failed state for all, one, or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetFailedCmd).Standalone()

	rootCmd.AddCommand(resetFailedCmd)

	carapace.Gen(resetFailedCmd).PositionalAnyCompletion(
		action.ActionUnits(resetFailedCmd).FilterArgs(),
	)
}
