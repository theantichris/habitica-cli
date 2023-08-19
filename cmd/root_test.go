package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_versionCMD(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	versionCmd.Execute()

	w.Close()
	os.Stdout = old

	out := make([]byte, 100)
	n, _ := r.Read(out)

	expected := "hcli v0"
	actual := strings.TrimSpace(string(out[:n]))

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
