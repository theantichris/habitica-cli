package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_Root(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("test Execute", func(t *testing.T) {
		Execute()

		out := make([]byte, 100)
		n, _ := r.Read(out)

		expected := "hcli v0"
		actual := strings.TrimSpace(string(out[:n]))

		if actual != expected {
			t.Errorf("expected '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("test version", func(t *testing.T) {
		versionCmd.Execute()

		out := make([]byte, 100)
		n, _ := r.Read(out)

		expected := "hcli v0"
		actual := strings.TrimSpace(string(out[:n]))

		if actual != expected {
			t.Errorf("expected '%s' but got '%s'", expected, actual)
		}
	})

	w.Close()
	os.Stdout = old
}
