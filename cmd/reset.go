/*
Copyright © 2026 Charles David Jorge <daviddevlife@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"InuSDK/internal/setup"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset InuSDK configuration",
	Long:  `Re-runs the setup wizard and resets configuration. Installed SDKs are not affected`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setup.Run(true); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
