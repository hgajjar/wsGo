package cmd

import (
	"bytes"
	"github.com/hgajjar/wsGo/types"
	"io/ioutil"
	"testing"
)

func TestBuildCommand(t *testing.T) {
	flagValue := "there"

	flags := []*types.Flag{
		types.NewFlag("name", "", "", nil),
	}
	cmdType := types.NewCommand("hello", "", flags, "echo \"Hello <name>\";")

	cmd := buildCommand(cmdType)
	buffer := bytes.NewBufferString("")
	cmd.SetOut(buffer)
	cmd.SetArgs([]string{"--name", flagValue})
	cmd.Execute()

	out, err := ioutil.ReadAll(buffer)

	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := "Hello " + flagValue

	if string(out) != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, string(out))
	}
}
