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

	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func Test_versionCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	versionCmd.SetOut(buf)

	versionCmd.Execute()

	expected := versionString
	actual := strings.Trim(versionString, " ")

	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}
