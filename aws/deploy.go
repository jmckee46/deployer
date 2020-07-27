package awsfuncs

import (
	"fmt"

	envvars "github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
)

func Deploy() flaw.Flaw {
	err := envvars.ValidateDeploymentEnvVars()
	if err != nil {
		fmt.Println(err.String())
	}

	return nil
}
