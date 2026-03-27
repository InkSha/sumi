package plugins

import (
	"fmt"
	"os"
	"path"

	"github.com/InkShaStudio/go-command"
	"github.com/inksha/sumi/internal/utils/common"
	"github.com/inksha/sumi/internal/utils/ufs"
)

func uninstall() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("plugin name to uninstall")

	cmd := command.NewCommand("uninstall").
		ChangeDescription("Uninstall a plugin").
		AddArgs(name).
		RegisterHandler(func(cmd *command.SCommand) {
			pluginDir := path.Join(cfg.PluginDir, name.Value)

			if !ufs.Exists(pluginDir) {
				common.Exit(fmt.Sprintf("Plugin %s is not installed", name.Value))
			}

			installedVersion := getInstalledVersion(name.Value)
			if installedVersion != "" {
				fmt.Printf("Uninstalling %s@%s...\n", name.Value, installedVersion)
			} else {
				fmt.Printf("Uninstalling %s...\n", name.Value)
			}

			if err := os.RemoveAll(pluginDir); err != nil {
				common.Exit(fmt.Sprintf("Failed to uninstall plugin: %s", err.Error()))
			}

			fmt.Printf("Plugin %s uninstalled successfully\n", name.Value)
		})

	return cmd
}
