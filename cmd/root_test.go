package cmd

import (
	"os"
	"strings"
	"testing"
)

var old, w *os.File

func Test_Execute(t *testing.T) {
	r := setUp()
	defer cleanUp()

	Execute()

	out := make([]byte, 100)
	n, _ := r.Read(out)

	expected := "hcli v0"
	actual := strings.TrimSpace(string(out[:n]))

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func Test_versionCmd(t *testing.T) {
	r := setUp()
	defer cleanUp()

	versionCmd.Execute()

	out := make([]byte, 100)
	n, _ := r.Read(out)

	expected := "hcli v0"
	actual := strings.TrimSpace(string(out[:n]))

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func setUp() *os.File {
	old = os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	return r
}

func cleanUp() {
	w.Close()
	os.Stdout = old
}
