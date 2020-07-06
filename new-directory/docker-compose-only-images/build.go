package dockerComposeOnlyImages

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/image"
)

// Build builds everything in docker-compose-only-images
func Build() flaw.Flaw {
	fmt.Println("docker-compose-only-images are building...")

	err := image.BuildWithTLS("docker-compose-only-images/load-balancer")
	if err != nil {
		return err
	}
	fmt.Println("")

	return nil
}
