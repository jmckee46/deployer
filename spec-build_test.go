package main

import (
	"fmt"
	"testing"

	"github.com/jmckee46/deployer/aws"
)

func TestCurrentSha(t *testing.T) {
	// err := awsfuncs.GetTlsFilesFromS3()
	state := awsfuncs.NewState()
	state.RenderedTemplateLocal = "artifacts/test-branch/completeStack"

	err := awsfuncs.ValidateTargetGroupNames(state)
	if err != nil {
		fmt.Println(err.String())
	}

	//
	//
	// fmt.Println("calling base unit test")
	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}

// "artifacts/test-bra-e3b0/completeStack"
