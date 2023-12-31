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
		fmt.Println("Hello, world!")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of hcli",
	Long:  "All software has versions, this is ours.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hcli v0")
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
