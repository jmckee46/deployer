package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// CreateMasterStack creates the initial master stack
func CreateMasterStack(state *State) flaw.Flaw {
	fmt.Println("  creating master stack...")

	// initialize cidr variables
	flawErr := InitializeCidrVariables(state)
	if flawErr != nil {
		return flawErr
	}

	// encrypt and store stack parameters in s3
	flawErr = StackParameters(state)
	if flawErr != nil {
		return flawErr
	}
	flawErr = PutStackParametersInS3(state)
	if flawErr != nil {
		return flawErr
	}

	fmt.Println("stack parameters:", state.StackParametersAll)
	// create master stack
	flawErr = CreateStack(state)
	if flawErr != nil {
		return flawErr
	}

	return nil
}
