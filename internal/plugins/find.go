package plugins

import (
	"fmt"
	"path"
	"runtime"

	"github.com/inksha/sumi/internal/utils/ufs"
)

func FindPlugin(name string) (string, error) {

	pluginDir := path.Join(cfg.PluginDir, name)

	if !ufs.Exists(pluginDir) {
		return "", fmt.Errorf("not found plugin %s", name)
	}

	system := runtime.GOOS
	arch := runtime.GOARCH

	executer := system + "-" + arch

	if system == "windows" {
		executer += ".exe"
	}

	executerPath := path.Join(pluginDir, executer)

	return executerPath, nil
}
