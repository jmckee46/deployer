package main

import (
	"os"

	"github.com/jmckee46/deployer/containers"
	"github.com/jmckee46/deployer/images"
)

func main() {

	switch os.Args[1] {
	case "build":
		build()
	case "crank":
		crank()
	case "init":
		initNewDirectory()
	case "start":
		start()
	case "stop":
		stop()
	case "test":
		test()
	case "prune":
		images.Prune()
		containers.Prune()
	default:
		os.Exit(1)
	}
}
