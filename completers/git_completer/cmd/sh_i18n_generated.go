package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sh_i18nCmd = &cobra.Command{
	Use:     "sh-i18n",
	Short:   "Git's i18n setup code for shell scripts",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(sh_i18nCmd).Standalone()

	rootCmd.AddCommand(sh_i18nCmd)
}
