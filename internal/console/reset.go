package console

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the bucket",
	Run: func(cmd *cobra.Command, _ []string) {
		login, _ := cmd.Flags().GetString("login")
		ip, _ := cmd.Flags().GetString("ip")

		err := accessClient.ResetBucket(context.Background(), login, ip)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Bucket reset successfully cleared")
	},
}

func init() {
	resetCmd.Flags().StringP("login", "l", "", "Login to reset the bucket")
	resetCmd.Flags().StringP("ip", "i", "", "IP address to reset the bucket")
	err := resetCmd.MarkFlagRequired("login")
	if err != nil {
		return
	}
	err = resetCmd.MarkFlagRequired("ip")
	if err != nil {
		return
	}

	rootCmd.AddCommand(resetCmd)
}
