package sumi_template

import (
	"os"
	"path"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/inksha/sumi/internal/utils/common"
	"github.com/inksha/sumi/internal/utils/ufs"
)

var templateDir = path.Join(os.Getenv("USERPROFILE"), ".sumi", "templates")

func register() *command.SCommand {
	n := command.NewCommandArg[string]("name").ChangeDescription("Template name").ChangeValue("")
	f := command.NewCommandFlag[[]string]("files").ChangeDescription("Template path").ChangeValue([]string{})
	d := command.NewCommandFlag[[]string]("dir").ChangeDescription("Template directory").ChangeValue([]string{})
	e := command.NewCommandFlag[[]string]("exclude").ChangeDescription("Exclude paths")
	c := command.NewCommandFlag[string]("comment").ChangeDescription("Template comment").ChangeValue("")

	cmd := command.
		NewCommand("register").
		ChangeDescription("Register new template.").
		AddArgs(n).
		AddFlags(f, d, e, c).
		RegisterHandler(func(cmd *command.SCommand) {
			if n.Value == "" {
				common.Exit("Template name is required.")
			}

			if strings.Contains(n.Value, ".") {
				common.Exit("Template name cannot contain dots.")
			}

			if len(f.Value) == 0 && len(d.Value) == 0 {
				common.Exit("At least one template path or directory is required.")
			}

			dir := path.Join(templateDir, n.Value)

			err := ufs.MkDir(dir, true)

			if err != nil {
				if os.IsExist(err) {
					common.Exit("Template name already exists.")
				} else {
					common.Exit("Failed to create template directory.")
				}
			}

			paths := []string{}

			for _, p := range f.Value {
				info, err := os.Stat(p)

				if err != nil {
					common.Exit("Failed to stat template path.")
				}

				if info.IsDir() {
					paths = append(paths, ufs.WalkListDir(p, e.Value)...)
				} else {
					paths = append(paths, p)
				}
			}

			for _, p := range d.Value {
				paths = append(paths, ufs.WalkListDir(p, e.Value)...)
			}

			if len(paths) == 0 {
				common.Exit("No template files found.")
			}

			for _, p := range paths {
				target := path.Join(dir, p)

				_, err := ufs.Copy(p, target)

				if err != nil {
					common.Exit("Failed to copy template file.")
				}
			}

			if c.Value != "" {
				commentFile := path.Join(templateDir, "."+n.Value)
				err := ufs.WriteFile(commentFile, c.Value)

				if err != nil {
					common.Exit("Failed to write comment file.")
				}
			}

		})

	return cmd
}
