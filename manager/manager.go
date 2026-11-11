package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"configs-manager/utils"

	// NOTE: rename to git-module-manager? git-pkg-manager? git-submodule-manager?
	modulesManager "github.com/SergeyDarn/easy-modules-manager/manager"
	modulesUtils "github.com/SergeyDarn/easy-modules-manager/utils"
	"github.com/charmbracelet/log"
)

func linkConfig(config utils.Config, cwd string) {
	fmt.Println("🪚 config:", config)

	_, err := os.Lstat(config.Dist)
	isDistExists := !os.IsNotExist(err)

	if err != nil && isDistExists {
		errText := fmt.Sprintf("skipping %s, err: %s", config.Src, err.Error())
		modulesUtils.LogError(errText)
		return
	}

	// NOTE: ? add warning or an input for user in case he doesn't want to forcefully replace config?

	if isDistExists {
		err = os.RemoveAll(config.Dist)
		if err != nil {
			// TODO: add logf functions to my utils and refactor calls like this
			errText := fmt.Sprintf("error while deleting symlink dist %s, err: %s", config.Dist, err.Error())
			modulesUtils.LogError(errText)
			return
		}
	}

	srcFullPath := filepath.Join(
		cwd,
		modulesUtils.GetModulesDir(),
		config.Src,
	)
	fmt.Println("🪚 srcFullPath:", srcFullPath)
	fmt.Println("🪚 config.Dist:", config.Dist)
	fmt.Println()
	utils.Symlink(srcFullPath, config.Dist)
}

func Init(shouldCloneRepos bool) {
	if shouldCloneRepos {
		modulesManager.Init(log.DebugLevel)
	} else {
		// Required to set modules directory
		modulesUtils.InitEnv()
	}

	cwd, err := os.Getwd()
	modulesUtils.CheckError(err, "couldn't get current working directory")

	// TODO: add conccurency here
	configJson := utils.ReadConfigJson()
	for _, config := range configJson.Configs {
		linkConfig(config, cwd)
	}
}
