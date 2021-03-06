package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/docker"
	"github.com/jmckee46/deployer/flaw"
)

// PushDockerImages prepares and pushes docker images to ecr
func PushDockerImages(state *State) flaw.Flaw {
	fmt.Println("  pushing docker images to ECR...")

	// log docker into ecr registry
	err := docker.Login(state.GetDockerRegistry())
	if err != nil {
		return err
	}

	// create ecr repositories
	err = CreateEcrRepositories(state)
	if err != nil {
		return err
	}

	// push docker images
	err = docker.Push(state.ImagesToDeploy, state.DockerRegistry)
	if err != nil {
		return err
	}

	return nil
}
