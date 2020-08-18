package awsfuncs

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// TransitionStack
func TransitionStack(state *State) flaw.Flaw {
	fmt.Println("Transition stack...")
	if NotOKToDeploy() {
		fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
		return nil
	}

	fmt.Println("optionally transitioning stack...")
	// get stack parameters
	flawErr := StackParameters(state)
	if flawErr != nil {
		return flawErr
	}

	if os.Getenv("DE_STACK_NAME") == "master" {
		// encrypt and store stack parameters in s3
		flawErr := PutStackParametersInS3(state)
		if flawErr != nil {
			return flawErr
		}

		if MasterStackExists(state) {
			// update master via change set
		} else {
			// create master stack
			flawErr := CreateMasterStack(state)
			if flawErr != nil {
				return flawErr
			}
		}
	} else {
		// 	create branch master
		// 	update via change-set
	}

	return nil
}
