package docker

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/myAppTest/docker-compose-only-images"
	"github.com/jmckee46/myAppTest/docker-compose-only-images/load-balancer"
	"github.com/jmckee46/myAppTest/images-to-deploy"
)

// Build runs a complete build of everything
func Build() flaw.Flaw {
	err := preBuild()
	if err != nil {
		return err
	}

	err = dockerComposeOnlyImages.Build()
	if err != nil {
		return err
	}

	err = imagesToDeploy.Build()
	if err != nil {
		return err
	}

	err = postBuild()
	if err != nil {
		return err
	}

	return nil
}
func preBuild() flaw.Flaw {
	fmt.Println("pre-build started...")

	err := loadBalancer.CreateCert()
	if err != nil {
		return err
	}
	fmt.Println("pre-build completed...")
	fmt.Println("")

	return nil
}

func postBuild() flaw.Flaw {
	fmt.Println("post-build started...")

	err := loadBalancer.DeleteCert()
	if err != nil {
		return err
	}
	fmt.Println("post-build completed...")
	fmt.Println("")

	return nil
}
