package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hcli",
	Short: "hcli is a CLI for Habitica.",
	Long:  "Learn more about Habitica at https://habitica.com.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here.
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
