package dockerCompose

import (
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Stop stops all containers
func Stop() flaw.Flaw {
	fmt.Println("stopping containers...")

	cmd := exec.Command(
		"docker-compose",
		"down",
		"--remove-orphans",
	)
	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
