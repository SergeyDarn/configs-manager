// Package manager provides main application logic
package manager

import (
	"os"
	"path/filepath"
	"sync"

	"configs-manager/utils"

	reposManager "github.com/SergeyDarn/git-repos-manager/manager"
	reposUtils "github.com/SergeyDarn/git-repos-manager/utils"
	"github.com/charmbracelet/log"
)

func linkConfig(config utils.Config, cwd string) {
	_, err := os.Lstat(config.Dist)
	isDistExists := !os.IsNotExist(err)

	if err != nil && isDistExists {
		reposUtils.LogErrorf("skipping %s, err: %s", config.Src, err.Error())
		return
	}

	if isDistExists {
		err = os.RemoveAll(config.Dist)
		if err != nil {
			reposUtils.LogErrorf("error while deleting symlink; dist %s, err: %s", config.Dist, err.Error())
			return
		}
	}

	srcFullPath := filepath.Join(
		cwd,
		reposUtils.GetReposDir(),
		config.Src,
	)
	utils.Symlink(srcFullPath, config.Dist)
}

func Init(shouldCloneRepos bool) {
	if shouldCloneRepos {
		reposManager.Init(log.DebugLevel)
	} else {
		// Required to set repos directory
		reposUtils.InitEnv()
	}

	cwd, err := os.Getwd()
	reposUtils.CheckError(err, "couldn't get current working directory")

	var waitGroup sync.WaitGroup

	configJSON := utils.ReadConfigJSON()
	for _, config := range configJSON.Configs {
		waitGroup.Add(1)

		go func() {
			linkConfig(config, cwd)
			defer waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
