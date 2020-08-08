package awsfuncs

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// OptionallyTransitionStack
func OptionallyTransitionStack(state *state) flaw.Flaw {
	if NotOKToDeploy() {
		fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
		return nil
	}

	fmt.Println("optionally transitioning stack...")

	if os.Getenv("DE_STACK_NAME") == "master" {
		// encrypt and store stack parameters in s3
		//   convert to json via json.MarshalIndent or Marshal output is []byte
		//   encrypt the []byte
		//   write to local file ioutil.WriteFile
		//   store in s3 via put-file-in-s3

		if MasterStackExists(state) {
			// update master via change set
		} else {
			// create master stack
		}
	} else {
		// 	create branch master
		// 	update via change-set
	}

	return nil
}
