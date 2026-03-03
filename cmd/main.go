package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/InkShaStudio/go-command"
	"github.com/inksha/sumi/internal/plugins"
	"github.com/inksha/sumi/internal/utils/common"
)

func Execute() {
	localPlugins := plugins.ListPlugin()

	subCommands := []*command.SCommand{
		plugins.RegisterCommand(),
	}

	for _, details := range localPlugins {
		args := []command.ICommandArgValue{}
		flags := []command.ICommandFlagValue{}

		for _, arg := range details.Manifest.Args {
			args = append(args,
				command.
					NewCommandArg[string](arg.Name).
					ChangeDescription(arg.Description).
					ChangeValue(arg.Value),
			)
		}

		for _, flag := range details.Manifest.Flags {
			flags = append(flags,
				command.
					NewCommandFlag[string](flag.Name).
					ChangeDescription(flag.Description).
					ChangeValue(flag.Value),
			)
		}

		cmd := command.
			NewCommand(details.Name).
			ChangeDescription(details.Manifest.Description).
			AddArgs(args...).
			AddFlags(flags...).
			RegisterHandler(func(cmd *command.SCommand) {
				restArgs := []string{}

				for _, arg := range cmd.Args {
					current := arg.GetValue().(*string)
					restArgs = append(restArgs, *current)
				}

				for _, flag := range cmd.Flags {
					current := flag.GetValue().(*string)
					if *current == "" {
						continue
					}

					restArgs = append(restArgs, fmt.Sprintf("--%s=%s", flag.GetName(), *current))
				}

				execute := exec.Command(details.Execute, restArgs...)

				output, err := execute.CombinedOutput()

				if err != nil {
					common.Exit(err.Error())
				}

				fmt.Println(string(output))
			})

		subCommands = append(subCommands, cmd)
	}

	sumi := command.NewCommand("sumi").
		ChangeDescription("Always want to see summer in you eyes.").
		RegisterHandler(func(cmd *command.SCommand) {
			fmt.Println("Hello, Sumi!")
		}).
		AddSubCommand(
			subCommands...,
		)

	cmd := command.RegisterCommand(sumi)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
