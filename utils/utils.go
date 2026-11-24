// Package utils provides general utilities
package utils

import (
	"os"

	reposUtils "github.com/SergeyDarn/git-repos-manager/utils"
)

type Config struct {
	Src  string
	Dist string
}

type ConfigJSON struct {
	Configs []Config
}

const configFile = "configs.json"

func ReadConfigJSON() ConfigJSON {
	return reposUtils.ReadJSON[ConfigJSON](configFile)
}

func Symlink(src string, dist string) {
	err := os.Symlink(src, dist)
	if err == nil {
		reposUtils.LogSuccessf("symlink to %s successfully created", dist)
		return
	}

	reposUtils.LogErrorf(
		"couldn't create symlink for config %s %s, err: %s",
		src,
		dist,
		err.Error(),
	)
}
