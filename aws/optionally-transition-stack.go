package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// RenderStackTemplate assembles the various stack templates into one
func OptionallyTransictionStack(state *state) flaw.Flaw {
	if NotOKToDeploy() {
		fmt.Println("not deploying to aws: not on master and HEAD comment does not end [deploy]")
		return nil
	}

	fmt.Println("optionally transitioning stack...")

	return nil
}