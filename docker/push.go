package docker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Push pushes all images to deploy to ecr
func Push(imagesToDeploy []string, dockerRegistry string) flaw.Flaw {
	fmt.Println("pushing images...")
	for _, image := range imagesToDeploy {
		fmt.Printf("%s...\n", image)
		location := dockerRegistry + "/" + os.Getenv("DE_STACK_NAME")
		location += "-" + image + ":" + os.Getenv("DE_GIT_SHA")
		fmt.Println("location:", location)

		localImage := "myapptest/images-to-deploy/" + image

		// tag image
		cmd := exec.Command(
			"docker",
			"tag",
			localImage,
			location,
		)
		tagResult, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("err:", err)
			return flaw.From(err)
		}
		fmt.Println("tagResult:", string(tagResult))

		// push image into registry
		cmd = exec.Command(
			"docker",
			"push",
			location,
		)
		pushResult, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("err:", err)
			return flaw.From(err)
		}
		fmt.Println("pushResult:", string(pushResult))
	}

	return nil
}
