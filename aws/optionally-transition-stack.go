package awsfuncs

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// OptionallyTransitionStack either creates a new stack, or if one exists it updates via change set.
// If the stack is a branch it will initially create a copy of the master stack in the branch
// then update via change set. A copy of master is only created on the initial branch deploy.
// NOTE: the first deploy would test updating the master stack, but subsequent
// deploys would not. Before submitting a pull request the branch stack
// should be deleted and the process started fresh to test updating master...
func OptionallyTransitionStack(state *State) flaw.Flaw {
	if NotOKToDeploy() {
		fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
		return nil
	}

	fmt.Println("optionally transitioning stack...")
	// push images to ecr
	flawErr := PushDockerImages(state)
	if flawErr != nil {
		return flawErr
	}

	if os.Getenv("DE_STACK_NAME") == "master" {
		// encrypt and store stack parameters in s3
		flawErr := PutStackParametersInS3(state)
		if flawErr != nil {
			return flawErr
		}
		// describe master stack to see if it exists
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
		// describe branch stack to see if it exists
		// if branch stack does not exist
		//   then get master parameters (StackParametersStackState) from s3
		//        combine master parameters (StackParametersStackState) with current StackParametersPasswords
		//        get master sha from a describe of master stack
		//        create branch master using combined master parameters and master template
		//   else update branch stack via change set
	}

	// aws/put-load-balancer-dns
	// aws/run-migrator
	// aws/prune-certificates-from-acm

	return nil
}
