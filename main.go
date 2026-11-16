package main

import (
	"flag"

	"configs-manager/manager"

	"github.com/SergeyDarn/easy-modules-manager/modules"
	modulesUtils "github.com/SergeyDarn/easy-modules-manager/utils"
)

func main() {
	shouldCloneRepos := flag.Bool(
		"clone",
		false,
		"Set to true on first usage to clone all the configs' repositories",
	)
	shouldPullRepos := flag.Bool(
		"pull",
		false,
		"Perform git pull for all your configs' repositories",
	)
	flag.Parse()

	if *shouldPullRepos {
		modulesUtils.InitEnv()
		modules.PullModules()
		return
	}

	manager.Init(*shouldCloneRepos)
}
