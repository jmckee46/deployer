package awsfuncs

import (
	envvars "github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
)

// Deploy
func Deploy() flaw.Flaw {
	// validate env vars
	err := envvars.ValidateDeploymentEnvVars()
	if err != nil {
		return err
	}

	// TODO: NEED TO WRITE CREATE STACK BEFORE I CAN DELETE...
	// // consider delete stack
	// err = considerDeleteStack()
	// if err != nil {
	// 	return err
	// }

	// prepare the stack template
	err = PrepareStackTemplate()
	if err != nil {
		return err
	}

	return nil
}
