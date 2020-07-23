package tlsDeployer

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

func manageFiles() flaw.Flaw {
	fmt.Println("managing tls files...")

	state := newState()

	if tlsFilesInS3Fresh(state) {
		err := getTLSFilesFromS3(state)
		if err != nil {
			return err
		}
	}

	return nil
}
