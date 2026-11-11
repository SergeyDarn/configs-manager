package utils

import (
	"fmt"
	"os"

	"github.com/SergeyDarn/easy-modules-manager/utils"
)

type Config struct {
	Src  string
	Dist string
}

type ConfigJson struct {
	Configs []Config
}

var configFile = "configs.json"

func ReadConfigJson() ConfigJson {
	return utils.ReadJson[ConfigJson](configFile)
}

func Symlink(src string, dist string) {
	err := os.Symlink(src, dist)
	if err == nil {
		return
	}

	utils.LogError(
		fmt.Sprintf(
			"couldn't create symlink for config %s %s, err: %s",
			src,
			dist,
			err.Error(),
		),
	)
}
