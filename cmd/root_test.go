package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func Test_rootCmd(t *testing.T) {
	root := &cobra.Command{Use: "root", Run: rootCmd.Run}

	buffer := new(bytes.Buffer)
	root.SetOut(buffer)

	root.Execute()

	expected := welcomeString
	actual := strings.TrimSpace(welcomeString)

	if expected != actual {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func Test_versionCmd(t *testing.T) {
	version := &cobra.Command{Use: "version", Run: versionCmd.Run}

	buffer := new(bytes.Buffer)
	version.SetOut(buffer)

	version.Execute()

	expected := versionString
	actual := strings.TrimSpace(versionString)

	if expected != actual {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
