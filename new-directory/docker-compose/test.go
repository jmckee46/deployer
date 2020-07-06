package dockerCompose

import (
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Test runs http-specs on docker containers
func Test() flaw.Flaw {
	fmt.Println("deployer testing...")

	cmd := exec.Command(
		"docker-compose",
		"run",
		"http-specs",
	)
	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
