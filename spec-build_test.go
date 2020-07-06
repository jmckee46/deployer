package main

import (
	"fmt"
	"testing"

	"github.com/jmckee46/deployer/containers"
)

func TestCurrentSha(t *testing.T) {
	err := containers.Prune()
	// err := gofuncs.Build("images-to-deploy/health-check")

	if err != nil {
		fmt.Println(err.String())
	}

	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
