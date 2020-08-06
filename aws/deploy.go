package awsfuncs

import (
	"fmt"

	envvars "github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
)

// Deploy deploys docker images to AWS
func Deploy() flaw.Flaw {
	fmt.Println("Deploying to AWS")

	state := NewState()

	// validate env vars
	err := envvars.ValidateDeploymentEnvVars()
	if err != nil {
		return err
	}

	// TODO: NEED TO WRITE CREATE STACK BEFORE I CAN DELETE...
	// // consider delete stack
	// err = considerDeleteStack(state)
	// if err != nil {
	// 	return err
	// }

	// if NotOKToDeploy() {
	// 	fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
	// 	return nil
	// }

	// prepare the stack template
	err = PrepareStackTemplate(state)
	if err != nil {
		return err
	}

	// optionially transition stack
	err = OptionallyTransitionStack(state)
	if err != nil {
		return err
	}

	return nil
}
