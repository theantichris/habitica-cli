package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theantichris/habitica-cli/cmd/habitica"
)

const welcomeString = "Thank you for using hcli."
const versionString = "hcli v0"

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
	rootCmd.PersistentFlags().String("apiToken", "", "You API token")

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(habitica.StatusCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
