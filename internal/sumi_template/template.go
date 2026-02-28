package sumi_template

import (
	"github.com/InkShaStudio/go-command"
)

func RegisterCommand() *command.SCommand {
	cmd := command.
		NewCommand("template").
		ChangeDescription("Template commands.").
		RegisterHandler(func(cmd *command.SCommand) {
			listTemplate()
		}).
		AddSubCommand(
			register(),
			list(),
			useTemplate(),
		)

	return cmd
}
