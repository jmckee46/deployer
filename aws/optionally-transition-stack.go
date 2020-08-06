package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// OptionallyTransitionStack
func OptionallyTransitionStack(state *state) flaw.Flaw {
	if NotOKToDeploy() {
		fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
		return nil
	}

	fmt.Println("optionally transitioning stack...")

	// if branch = master
	// 	check master exists, if not then create it
	// 	otherwise update via change-set

	// if branch = branch-StackParameters
	// 	create branch master
	// 	update via change-set

	if MasterStackExists(state) {

	}

	return nil
}
