package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number and exit (also --version)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)
}
