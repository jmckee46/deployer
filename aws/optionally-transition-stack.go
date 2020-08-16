package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// OptionallyTransitionStack
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

	// if os.Getenv("DE_STACK_NAME") == "master" {
	// 	// encrypt and store stack parameters in s3
	// 	flawErr := PutStackParametersInS3(state)
	// 	if flawErr != nil {
	// 		return flawErr
	// 	}
	// }

	return nil
}
