package deployLaptop

import (
	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/logger"
)

func DeployFromLaptop() {
	// initialize the environment
	flawErr := InitializeEnvironment()
	if flawErr != nil {
		logger.Panic("deploy-from-laptop", flawErr)
	}

	// deploy
	flawErr = awsfuncs.Deploy()
	if flawErr != nil {
		logger.Panic("deploy-from-laptop", flawErr)
	}

	// // clean up
	// flawErr = Cleanup()
	// if flawErr != nil {
	// 	logger.Panic("deploy-from-laptop", flawErr)
	// }
}
