package main

import (
	"flag"

	"configs-manager/manager"

	reposManager "github.com/SergeyDarn/git-repos-manager/manager"
	reposUtils "github.com/SergeyDarn/git-repos-manager/utils"
)

// TODO: Update README
// TODO: move configs.json with my setup configs into a separate repository
func main() {
	shouldCloneRepos := flag.Bool(
		"clone",
		false,
		"Set to true on first usage to clone all your repos from the config file",
	)
	shouldPullRepos := flag.Bool(
		"pull",
		false,
		"Perform git pull in all your repos",
	)
	flag.Parse()

	if *shouldPullRepos {
		reposUtils.InitEnv()
		reposManager.PullRepos()
		return
	}

	manager.Init(*shouldCloneRepos)
}
