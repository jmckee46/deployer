package imagesToDeploy

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/image"
	"github.com/jmckee46/myAppTest/images-to-deploy/http-specs/build-http-specs"
)

// Build builds everything in images-to-deploy
func Build() flaw.Flaw {
	fmt.Println("images-to-deploy are building...")

	// build http-specs
	err := buildHttpSpecs.Build()
	if err != nil {
		return err
	}

	// build health-check
	err = image.BuildGoWithTLS("images-to-deploy/health-check")
	if err != nil {
		return err
	}

	// build migrator
	err = BuildMigrator()
	if err != nil {
		return err
	}
	fmt.Println("")

	return nil
}
