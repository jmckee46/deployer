package main

import (
	"fmt"
	"os"

	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	// set up env-vars
	flawErr := initEnvVars()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// create aws config file
	flawErr = awsfuncs.CreateCliConfigFile()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

	// get go dependencies
	// flawErr = gofuncs.GetDependencies()
	// if flawErr != nil {
	// 	logger.Panic("travis-ci-install", flawErr)
	// }

	// create artifacts directory
	fmt.Println("creating artifacts directory")
	err := os.MkdirAll(os.Getenv("DE_ARTIFACTS_PATH"), 0755)
	if err != nil {
		logger.Panic("travis-ci-install", flaw.From(err))
	}

	// initialize certbot-env
	flawErr = initCertbotEnv()
	if flawErr != nil {
		logger.Panic("travis-ci-install", flawErr)
	}

}
