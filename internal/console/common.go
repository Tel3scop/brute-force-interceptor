package console

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// AddCommand is a common function for adding a subnet to a list (blacklist or whitelist).
func AddCommand(cmd *cobra.Command, addFunc func(context.Context, string) error, successMessage string) {
	subnet, _ := cmd.Flags().GetString("subnet")

	err := addFunc(context.Background(), subnet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(successMessage)
}

// RemoveCommand is a common function for removing a subnet from a list (blacklist or whitelist).
func RemoveCommand(cmd *cobra.Command, removeFunc func(context.Context, string) error, successMessage string) {
	subnet, _ := cmd.Flags().GetString("subnet")

	err := removeFunc(context.Background(), subnet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(successMessage)
}

// InitCommand initializes a command with a subnet flag and marks it as required.
func InitCommand(cmd *cobra.Command, flagName, shorthand, usage string) {
	cmd.Flags().StringP(flagName, shorthand, "", usage)
	err := cmd.MarkFlagRequired(flagName)
	if err != nil {
		return
	}
}
