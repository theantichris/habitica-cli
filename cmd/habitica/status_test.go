package habitica

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func Test_statusCmd(t *testing.T) {
	status := &cobra.Command{Use: "status", Run: statusCmd.Run}

	buffer := new(bytes.Buffer)
	status.SetOut(buffer)

	status.Execute()

	expected := statusOkString
	actual := strings.TrimSpace(buffer.String())

	if expected != actual {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
