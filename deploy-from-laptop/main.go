package main

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/tls"

	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	// create aws config file
	fmt.Println("creating aws config file...")
	flawErr := awsfuncs.CreateCliConfigFile()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// get go dependencies
	// flawErr = gofuncs.GetDependencies()
	// if flawErr != nil {
	// 	logger.Panic("travis-ci-install", flawErr)
	// }

	// create artifacts directory
	fmt.Println("creating artifacts directory...")
	err := os.MkdirAll(os.Getenv("DE_ARTIFACTS_PATH"), 0755)
	if err != nil {
		logger.Panic("travis-ci-install", flaw.From(err))
	}

	// manage tls files
	flawErr = tlsDeployer.ManageFiles()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// clean up
	flawErr = Cleanup()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}
}
