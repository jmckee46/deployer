package loadBalancer

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// DeleteCert deletes the load balancer certs
func DeleteCert() flaw.Flaw {
	fmt.Println("deleting load-balancer cert...")

	err := os.Remove("docker-compose-only-images/load-balancer/load-balancer.crt")
	if err != nil {
		return flaw.From(err)
	}

	err = os.Remove("docker-compose-only-images/load-balancer/load-balancer.key")
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
