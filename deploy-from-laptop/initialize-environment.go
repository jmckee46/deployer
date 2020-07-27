package deployLaptop

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/tls"

	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/flaw"
)

func InitializeEnvironment() flaw.Flaw {
	// create aws config file
	flawErr := awsfuncs.CreateCliConfigFile()
	if flawErr != nil {
		return flawErr
	}

	// get go dependencies
	// flawErr = gofuncs.GetDependencies()
	// if flawErr != nil {
	// return flawErr
	// }

	// create artifacts directory
	fmt.Println("creating artifacts directory...")
	err := os.MkdirAll(os.Getenv("DE_ARTIFACTS_PATH"), 0755)
	if err != nil {
		return flawErr
	}

	// manage tls files
	flawErr = tlsDeployer.ManageFiles()
	if flawErr != nil {
		return flawErr
	}

	return nil
}
