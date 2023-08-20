package habitica

import (
	"bytes"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

const authorId = "cf3a800e-4359-402c-90bc-e9d87ee78379"
const scriptName = "hcli"

const statusOkString = "status: ok"

// StatusCmd checks the status of the Habitica API, it returns 'status: ok'
// if it is up.
// hcli status --apiToken butts
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the Habitica API status.",
	Long:  "Checks the Habitica API status, returns 'status: ok' if it is up.",
	Run: func(cmd *cobra.Command, args []string) {
		apiTokenFlag := cmd.Flag("apiToken")

		if apiTokenFlag != nil {
			apiToken := apiTokenFlag.Value.String()

			url := "https://habitica.com/api/v3/status"
			data := []byte("")
			request, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
			if err != nil {
				cmd.PrintErrf("Error creating request: %t", err)
			}

			request.Header.Set("x-client", authorId+"-"+scriptName)
			request.Header.Set("x-api-user", authorId)
			request.Header.Set("x-api-key", apiToken)
			request.Header.Set("content-type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				cmd.PrintErrf("API request err: %t", err)
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				cmd.PrintErrf("Error reading body: %t", err)
			}

			cmd.Println(string(body))
		}

		// TODO: handle no API key
	},
}
