package dockerCompose

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/jmckee46/deployer/flaw"
)

// Start starts all docker containers
func Start() flaw.Flaw {
	fmt.Println("starting containers...")

	cmd := exec.Command(
		"docker-compose",
		"up",
		"--detach",
		"--force-recreate",
		"--remove-orphans",
		"health-check",
		"load-balancer",
		"migrations-postgres",
	)
	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return flaw.From(err)
	}

	// run migrations
	time.Sleep(10 * time.Second)

	fmt.Println("running migrations...")

	cmd = exec.Command(
		"docker-compose",
		"run",
		"migrator",
	)
	stdoutStderr, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		return flaw.From(err)
	}

	return nil
}
