package plugins

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/google/go-github/v84/github"
	"github.com/inksha/sumi/internal/utils/api"
	"github.com/inksha/sumi/internal/utils/common"
	"github.com/inksha/sumi/internal/utils/ufs"
	sumicc "github.com/summink/sumi-common-command"
)

func showLocalInfo(pluginName string) {
	pluginDir := path.Join(cfg.PluginDir, pluginName)
	manifestPath := path.Join(pluginDir, manifestFile)

	if !ufs.Exists(manifestPath) {
		common.Exit(fmt.Sprintf("Plugin %s is not installed", pluginName))
	}

	data, err := ufs.ReadFileByByte(manifestPath)
	if err != nil {
		common.Exit(fmt.Sprintf("Failed to read manifest: %s", err.Error()))
	}

	var manifest sumicc.PluginManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		common.Exit(fmt.Sprintf("Failed to parse manifest: %s", err.Error()))
	}

	fmt.Printf("Name:        %s\n", manifest.Name)
	fmt.Printf("Version:     %s\n", manifest.Version)
	fmt.Printf("Description: %s\n", manifest.Description)
	fmt.Printf("Author:      %s\n", manifest.Author)
	fmt.Printf("Location:    %s\n", pluginDir)
}

func showOnlineInfo(pluginName string) {
	repo := pluginPrefix + pluginName
	client := github.NewClient(nil)

	repoInfo, _, err := client.Repositories.Get(context.Background(), orgName, repo)
	if err != nil {
		common.Exit(fmt.Sprintf("Failed to get repository info: %s", err.Error()))
	}

	url := strings.ReplaceAll(repoInfo.GetContentsURL(), "{+path}", manifestFile)
	reps, err := api.Get(url)
	if err != nil {
		common.Exit(fmt.Sprintf("Failed to get manifest: %s", err.Error()))
	}

	downloadURL, ok := reps["download_url"].(string)
	if !ok {
		common.Exit("Invalid manifest URL")
	}

	rawManifest, err := api.GetRaw(downloadURL)
	if err != nil {
		common.Exit(fmt.Sprintf("Failed to download manifest: %s", err.Error()))
	}

	var manifest sumicc.PluginManifest
	if err := json.Unmarshal(rawManifest, &manifest); err != nil {
		common.Exit(fmt.Sprintf("Failed to parse manifest: %s", err.Error()))
	}

	releases, _, err := client.Repositories.ListReleases(context.Background(), orgName, repo, nil)
	latestVersion := "unknown"
	if err == nil && len(releases) > 0 {
		latestVersion = releases[0].GetTagName()
	}

	fmt.Printf("Name:           %s\n", manifest.Name)
	fmt.Printf("Latest Version: %s\n", latestVersion)
	fmt.Printf("Description:    %s\n", manifest.Description)
	fmt.Printf("Author:         %s\n", manifest.Author)
	fmt.Printf("Repository:     https://github.com/%s/%s\n", orgName, repo)

	installedVersion := getInstalledVersion(pluginName)
	if installedVersion != "" {
		fmt.Printf("Installed:      %s\n", installedVersion)
	}
}

func info() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("plugin name to show info")
	online := command.NewCommandFlag[bool]("online").ChangeDescription("get info from online")

	cmd := command.NewCommand("info").
		ChangeDescription("Show plugin information").
		AddArgs(name).
		AddFlags(online).
		RegisterHandler(func(cmd *command.SCommand) {
			if online.Value {
				showOnlineInfo(name.Value)
			} else {
				showLocalInfo(name.Value)
			}
		})

	return cmd
}
