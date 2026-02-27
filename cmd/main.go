package cmd

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/go-command"
)

func Execute() {
	cmd := command.RegisterCommand(
		command.NewCommand("sumi").
			ChangeDescription("Always want to see summer in you eyes.").
			RegisterHandler(func(cmd *command.SCommand) {
				fmt.Println("Hello, Sumi!")
			}),
	)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
