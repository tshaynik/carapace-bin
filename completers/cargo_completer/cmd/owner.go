package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Manage the owners of a crate on the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ownerCmd).Standalone()

	ownerCmd.Flags().StringSliceP("add", "a", []string{}, "Name of a user or team to invite as an owner")
	ownerCmd.Flags().BoolP("help", "h", false, "Print help")
	ownerCmd.Flags().String("index", "", "Registry index to modify owners for")
	ownerCmd.Flags().BoolP("list", "l", false, "List owners of a crate")
	ownerCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	ownerCmd.Flags().String("registry", "", "Registry to use")
	ownerCmd.Flags().StringSliceP("remove", "r", []string{}, "Name of a user or team to remove as an owner")
	ownerCmd.Flags().String("token", "", "API token to use when authenticating")
	rootCmd.AddCommand(ownerCmd)
}
