package console

import (
	"github.com/spf13/cobra"
)

var (
	whitelistAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add to whitelist",
		Run: func(cmd *cobra.Command, _ []string) {
			AddCommand(cmd, accessClient.AddToWhitelist, "Added to whitelist successfully")
		},
	}

	whitelistRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove from whitelist",
		Run: func(cmd *cobra.Command, _ []string) {
			RemoveCommand(cmd, accessClient.RemoveFromWhitelist, "Removed from whitelist successfully")
		},
	}
)

func init() {
	InitCommand(whitelistAddCmd, "subnet", "s", "Subnet to whitelist")
	InitCommand(whitelistRemoveCmd, "subnet", "s", "Subnet to whitelist")

	whitelistCmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Manage whitelist",
	}

	whitelistCmd.AddCommand(whitelistAddCmd)
	whitelistCmd.AddCommand(whitelistRemoveCmd)

	rootCmd.AddCommand(whitelistCmd)
}
