package console

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	whitelistAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add to whitelist",
		Run: func(cmd *cobra.Command, _ []string) {
			subnet, _ := cmd.Flags().GetString("subnet")

			err := accessClient.AddToWhitelist(context.Background(), subnet)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Added to whitelist successfully")
		},
	}

	whitelistRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove from whitelist",
		Run: func(cmd *cobra.Command, _ []string) {
			subnet, _ := cmd.Flags().GetString("subnet")

			err := accessClient.RemoveFromWhitelist(context.Background(), subnet)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Removed from whitelist successfully")
		},
	}
)

func init() {
	whitelistAddCmd.Flags().StringP("subnet", "s", "", "Subnet to whitelist")
	err := whitelistAddCmd.MarkFlagRequired("subnet")
	if err != nil {
		return
	}

	whitelistRemoveCmd.Flags().StringP("subnet", "s", "", "Subnet to whitelist")
	err = whitelistRemoveCmd.MarkFlagRequired("subnet")
	if err != nil {
		return
	}

	whitelistCmd := &cobra.Command{
		Use:   "whitelist",
		Short: "Manage whitelist",
	}

	whitelistCmd.AddCommand(whitelistAddCmd)
	whitelistCmd.AddCommand(whitelistRemoveCmd)

	rootCmd.AddCommand(whitelistCmd)
}
