package cmd

import (
	"github.com/spf13/cobra"
)

const welcomeString = "Thank you for using hcli."
const versionString = "hcli v0"

const authorId = "cf3a800e-4359-402c-90bc-e9d87ee78379"
const scriptName = "hcli"

var rootCmd = &cobra.Command{
	Use:   "hcli",
	Short: "hcli is a CLI for Habitica.",
	Long:  "Learn more about Habitica at https://habitica.com.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(welcomeString)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of hcli",
	Long:  "All software has versions, this is ours.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(versionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
