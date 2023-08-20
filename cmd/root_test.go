package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func Test_rootCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)

	rootCmd.Execute()

	expected := welcomeString
	actual := strings.Trim(welcomeString, " ")

	if actual != expected {
		t.Errorf("Expected: '%s', but got: '%s'", expected, actual)
	}
}
