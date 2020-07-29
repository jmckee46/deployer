package deployLaptop

import (
	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/tls"
)

func DeployFromLaptop() {
	// TravisCIInstall
	flawErr := TravisCIInstall()
	if flawErr != nil {
		logger.Panic("deploy-from-laptop", flawErr)
	}

	// manage tls files
	flawErr = tlsDeployer.ManageFiles()
	if flawErr != nil {
		logger.Panic("deploy-from-laptop", flawErr)
	}

	// // prepare to deploy - builds and tests docker images
	// flawErr = prepare()
	// if flawErr != nil {
	// 	logger.Panic("deploy-from-laptop", flawErr)
	// }

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
