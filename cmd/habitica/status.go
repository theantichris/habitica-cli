package habitica

import "github.com/spf13/cobra"

const authorId = "cf3a800e-4359-402c-90bc-e9d87ee78379"
const scriptName = "hcli"

const statusOkString = "status: ok"

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the Habitica API status.",
	Long:  "Checks the Habitica API status, returns 'status: ok' if it is up.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
