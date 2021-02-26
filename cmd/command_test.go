package cmd

import (
	"bytes"
	"github.com/hgajjar/wsGo/types/command"
	"io/ioutil"
	"testing"
)

func TestBuildCommand(t *testing.T) {
	flagValue := "there"

	flags := []*command.Flag{
		command.NewFlag("name", "", "", nil),
	}
	cmdType := command.NewCommand("hello", "", flags, "echo \"Hello <name>\";")

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
