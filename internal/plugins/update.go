package plugins

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/google/go-github/v84/github"
	"github.com/inksha/sumi/internal/utils/common"
)

func updatePlugin(pluginName string) {
	installedVersion := getInstalledVersion(pluginName)
	if installedVersion == "" {
		common.Exit(fmt.Sprintf("Plugin %s is not installed", pluginName))
	}

	client := github.NewClient(nil)
	repo := pluginPrefix + pluginName
	releases, _, err := client.Repositories.ListReleases(context.Background(), orgName, repo, nil)
	if err != nil {
		common.Exit(fmt.Sprintf("Failed to check updates: %s", err.Error()))
	}

	if len(releases) == 0 {
		common.Exit(fmt.Sprintf("No releases found for %s", pluginName))
	}

	latestVersion := releases[0].GetTagName()
	normalizedInstalled := installedVersion
	if !strings.HasPrefix(normalizedInstalled, "v") {
		normalizedInstalled = "v" + normalizedInstalled
	}

	if latestVersion == normalizedInstalled {
		fmt.Printf("Plugin %s@%s is already up to date\n", pluginName, installedVersion)
		return
	}

	fmt.Printf("Updating %s: %s -> %s\n", pluginName, installedVersion, latestVersion)
	downloadPlugin(repo, "", runtime.GOOS, runtime.GOARCH)
}

func updateAllPlugins() {
	plugins := findPlugin()
	if len(plugins) == 0 {
		fmt.Println("No plugins installed")
		return
	}

	for _, plugin := range plugins {
		pluginName := strings.ReplaceAll(plugin.Name, pluginPrefix, "")
		fmt.Printf("\nChecking %s...\n", pluginName)
		updatePlugin(pluginName)
	}
}

func update() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("plugin name to update (optional, update all if not specified)")

	cmd := command.NewCommand("update").
		ChangeDescription("Update plugin(s) to latest version").
		AddArgs(name).
		RegisterHandler(func(cmd *command.SCommand) {
			if name.Value != "" {
				updatePlugin(name.Value)
			} else {
				updateAllPlugins()
			}
		})

	return cmd
}
