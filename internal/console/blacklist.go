package console

import (
	"github.com/spf13/cobra"
)

var (
	blacklistAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add to blacklist",
		Run: func(cmd *cobra.Command, _ []string) {
			AddCommand(cmd, accessClient.AddToBlacklist, "Added to blacklist successfully")
		},
	}

	blacklistRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove from blacklist",
		Run: func(cmd *cobra.Command, _ []string) {
			RemoveCommand(cmd, accessClient.RemoveFromBlacklist, "Removed from blacklist successfully")
		},
	}
)

func init() {
	InitCommand(blacklistAddCmd, "subnet", "s", "Subnet to blacklist")
	InitCommand(blacklistRemoveCmd, "subnet", "s", "Subnet to blacklist")

	blacklistCmd := &cobra.Command{
		Use:   "blacklist",
		Short: "Manage blacklist",
	}

	blacklistCmd.AddCommand(blacklistAddCmd)
	blacklistCmd.AddCommand(blacklistRemoveCmd)

	rootCmd.AddCommand(blacklistCmd)
}
