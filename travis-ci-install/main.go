package main

import (
	"os"

	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/gofuncs"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	// create aws config file
	flawErr := awsfuncs.CreateCliConfigFile()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// get go dependencies
	flawErr = gofuncs.GetDependencies()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// create artifacts directory
	err := os.MkdirAll(os.Getenv("DE_ARTIFACTS_PATH"), 0755)
	if err != nil {
		logger.Panic("travis-ci-install", flaw.From(err))
	}
}
