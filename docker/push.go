package docker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Push pushes all images to deploy to ecr
func Push(imagesToDeploy []string, dockerRegistry string) flaw.Flaw {
	fmt.Println("    pushing images...")
	for _, image := range imagesToDeploy {
		fmt.Printf("      %s...\n", image)
		location := dockerRegistry + "/" + os.Getenv("DE_STACK_NAME")
		location += "-" + image + ":" + os.Getenv("DE_GIT_SHA")

		localImage := "myapptest/images-to-deploy/" + image

		// tag image
		cmd := exec.Command(
			"docker",
			"tag",
			localImage,
			location,
		)
		_, err := cmd.CombinedOutput()
		if err != nil {
			return flaw.From(err)
		}

		// push image into registry
		cmd = exec.Command(
			"docker",
			"push",
			location,
		)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return flaw.From(err)
		}
	}

	return nil
}
