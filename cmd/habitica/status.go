package habitica

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const authorId = "cf3a800e-4359-402c-90bc-e9d87ee78379"
const scriptName = "hcli"

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

			requester := HabiticaWebRequester{}
			body, err := requester.MakeRequest("GET", url, data, apiToken)
			if err != nil {
				cmd.PrintErr(err)
			}

			// {"success":true,"data":{"status":"up"},"appVersion":"5.1.4"}
			cmd.Println(string(body))
		}

		// TODO: handle no API key
	},
}

type WebRequester interface {
	MakeRequest(url string) (string, error)
}

type HabiticaWebRequester struct{}

func (r HabiticaWebRequester) MakeRequest(method string, url string, data []byte, apiToken string) (string, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return "", errors.Wrap(err, "Error creating request: %t")
	}

	request.Header.Set("x-client", authorId+"-"+scriptName)
	request.Header.Set("x-api-user", authorId)
	request.Header.Set("x-api-key", apiToken)
	request.Header.Set("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "API request err: %t")
	}
	defer response.Body.Close()

	fmt.Printf("Status code: %d\n", response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "Error reading body: %t")
	}

	return string(body), nil
}
