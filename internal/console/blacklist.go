package console

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	blacklistAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add to blacklist",
		Run: func(cmd *cobra.Command, _ []string) {
			subnet, _ := cmd.Flags().GetString("subnet")

			err := accessClient.AddToBlacklist(context.Background(), subnet)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Added to blacklist successfully")
		},
	}

	blacklistRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove from blacklist",
		Run: func(cmd *cobra.Command, _ []string) {
			subnet, _ := cmd.Flags().GetString("subnet")

			err := accessClient.RemoveFromBlacklist(context.Background(), subnet)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Removed from blacklist successfully")
		},
	}
)

func init() {
	blacklistAddCmd.Flags().StringP("subnet", "s", "", "Subnet to blacklist")
	err := blacklistAddCmd.MarkFlagRequired("subnet")
	if err != nil {
		return
	}

	blacklistRemoveCmd.Flags().StringP("subnet", "s", "", "Subnet to blacklist")
	err = blacklistRemoveCmd.MarkFlagRequired("subnet")
	if err != nil {
		return
	}

	blacklistCmd := &cobra.Command{
		Use:   "blacklist",
		Short: "Manage blacklist",
	}

	blacklistCmd.AddCommand(blacklistAddCmd)
	blacklistCmd.AddCommand(blacklistRemoveCmd)

	rootCmd.AddCommand(blacklistCmd)
}
