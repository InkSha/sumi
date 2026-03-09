package plugins

import (
	"context"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/google/go-github/v84/github"
	"github.com/inksha/sumi/internal/utils/api"
	"github.com/inksha/sumi/internal/utils/common"
	"github.com/inksha/sumi/internal/utils/ufs"
)

func exitGuard(err error) {
	if err == nil {
		return
	}

	common.Exit("install plugin failed: " + err.Error())
}

func downloadPlugin(repo string, version string, system string, arch string) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), orgName, repo, &github.ListOptions{})

	exitGuard(err)

	assetName := system + "-" + arch

	if system == "windows" {
		assetName += ".exe"
	}

topReleases:
	for _, release := range releases {
		for _, asset := range release.Assets {

			if strings.EqualFold(*asset.Name, assetName) || strings.HasSuffix(*asset.Name, assetName) {

				assetURL := asset.GetURL()
				assetInfo, _ := api.Get(assetURL)
				browserDownloadURL := assetInfo["browser_download_url"].(string)
				assetData, _ := api.GetRaw(browserDownloadURL)
				outputDir := path.Join(cfg.PluginDir, strings.ReplaceAll(repo, pluginPrefix, ""))
				output := path.Join(outputDir, assetName)

				exitGuard(ufs.MkDir(outputDir, true))
				exitGuard(ufs.WriteFileByByte(output, assetData))

				os.Chmod(output, 0755)

				println("install plugin "+repo+release.GetTagName()+" success to ", output)

				break topReleases
			}
		}
	}
}

func install() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("install plugin name")
	version := command.NewCommandArg[string]("version").ChangeDescription("install plugin version")

	cmd := command.NewCommand("install").
		ChangeDescription("Install a plugin").
		AddArgs(name, version).
		RegisterHandler(func(cmd *command.SCommand) {
			system := runtime.GOOS
			arch := runtime.GOARCH

			if ufs.Exists(path.Join(cfg.PluginDir, name.Value)) {
				println("Plugin " + name.Value + " already installed!")
				return
			}

			if version.Value != "" {
				println("Version parameter is not currently supported!")
			}

			downloadPlugin(pluginPrefix+name.Value, version.Value, system, arch)
		})

	return cmd
}
