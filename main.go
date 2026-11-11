package main

import (
	"flag"

	"configs-manager/manager"
)

func main() {
	shouldCloneRepos := flag.Bool(
		"clone-repos",
		false,
		"Set to true on first usage to clone all the git repos with the configs",
	)
	flag.Parse()

	manager.Init(*shouldCloneRepos)
}
