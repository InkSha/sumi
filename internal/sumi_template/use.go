package sumi_template

import (
	"fmt"
	"os"
	"path"

	"github.com/InkShaStudio/go-command"
	"github.com/inksha/sumi/internal/utils/common"
	"github.com/inksha/sumi/internal/utils/ufs"
)

func useTemplate() *command.SCommand {
	n := command.NewCommandArg[string]("name").ChangeDescription("use template name")

	t := command.NewCommandFlag[string]("target").ChangeDescription("target dir")
	r := command.NewCommandFlag[string]("rename").ChangeDescription("rename template generate dir")

	cmd := command.
		NewCommand("use").
		ChangeDescription("Use template.").
		AddArgs(n).
		AddFlags(r, t).
		RegisterHandler(func(cmd *command.SCommand) {
			if n.Value == "" {
				common.Exit("Template name is required.")
			}

			template := path.Join(templateDir, n.Value)

			if !ufs.Exists(template) {
				common.Exit("Template not exist.")
			}

			cwd, err := os.Getwd()

			if err != nil {
				common.Exit("Get current dir failed.")
			}

			target := cwd
			if t.Value != "" {
				target = t.Value
			}

			dir := path.Join(target, n.Value)
			if ufs.Exists(dir) {
				common.Exit("Target path already exist.")
			}

			ufs.MkDir(dir, true)

			ok, err := ufs.CopyAll(template, dir)
			if err != nil {
				common.Exit("Copy template failed.")
			}

			err = ufs.Rename(dir, path.Join(target, r.Value))
			if err != nil {
				println(err.Error())
				common.Exit("Rename template failed.")
			}

			if ok {
				fmt.Println("Use template success.")
			}
		})

	return cmd
}
