package cmd

import (
	"fmt"
	"github.com/hgajjar/wsGo/types"
	"github.com/hgajjar/wsGo/types/command"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func buildCommand(configType types.ConfigType) *cobra.Command {
	commandType := configType.(*command.Command)
	var cobraCommand = &cobra.Command{
		Use: commandType.Name(),
		Short: commandType.Usage(),
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			runScript := commandType.Script()

			for _, flag := range commandType.Flags() {
				runScript = strings.ReplaceAll(runScript, fmt.Sprintf("<%s>", flag.Name()), *flag.Value())
			}

			phpExecPath, err := exec.LookPath("php")

			if err != nil {
				return err
			}

			phpCmd := &exec.Cmd{
				Path: phpExecPath,
				Args: []string{ phpExecPath, "-r", runScript },
				Stdout: cobraCmd.OutOrStdout(),
				Stderr: cobraCmd.OutOrStderr(),
			}

			return phpCmd.Run()
		},
	}

	for _, flag := range commandType.Flags() {
		var flagValue string
		cobraCommand.Flags().StringVar(&flagValue, flag.Name(), flag.DefaultValue(), flag.Usage())
		flag.SetValue(&flagValue)
	}

	return cobraCommand
}