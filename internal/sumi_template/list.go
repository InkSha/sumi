package sumi_template

import (
	"fmt"
	"path"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/inksha/sumi/internal/utils/ufs"
)

func listTemplate() {
	paths := ufs.ListDir(templateDir)
	templates := []string{}
	comments := map[string]string{}
	maxNameLen := 0

	for _, p := range paths {
		name := path.Base(p)

		if name[0] == '.' {
			commentFile := path.Join(templateDir, name)
			comment, err := ufs.ReadFile(commentFile)

			if err != nil {
				continue
			}

			comments[name[1:]] = comment

		} else {
			templates = append(templates, name)
			maxNameLen = max(maxNameLen, len(name))
		}
	}

	for _, name := range templates {
		comment, ok := comments[name]

		if !ok {
			comment = ""
		}

		fmt.Printf("- %s%s : %s\n", name, strings.Repeat(" ", maxNameLen-len(name)), comment)
	}
}

func list() *command.SCommand {
	cmd := command.
		NewCommand("list").
		ChangeDescription("List all templates.").
		RegisterHandler(func(cmd *command.SCommand) {
			listTemplate()
		})

	return cmd
}
