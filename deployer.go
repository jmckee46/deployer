package main

import (
	"os"

	"github.com/jmckee46/deployer/containers"
	"github.com/jmckee46/deployer/deploy-from-laptop"
	"github.com/jmckee46/deployer/images"
)

func main() {

	switch os.Args[1] {
	case "build":
		build()
	case "deploy-from-laptop":
		deployLaptop.DeployFromLaptop()
	case "init":
		initNewDirectory()
	case "prepare":
		prepare()
	case "prune":
		images.Prune()
		containers.Prune()
	case "start":
		start()
	case "stop":
		stop()
	case "test":
		test()
	default:
		os.Exit(1)
	}
}
