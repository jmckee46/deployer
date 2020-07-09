package main

import (
	"fmt"
	"testing"

	"github.com/jmckee46/deployer/git"
)

func TestCurrentSha(t *testing.T) {
	branch, err := git.CurrentBranch()
	// err := gofuncs.Build("images-to-deploy/health-check")
	fmt.Println("branch:", branch)
	if err != nil {
		fmt.Println(err.String())
	}

	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
