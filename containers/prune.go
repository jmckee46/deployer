package containers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// Prune prunes exited docker images
func Prune() flaw.Flaw {
	fmt.Println("pruning containers...")

	// get container ids and clean them up
	cmd := exec.Command(
		"docker",
		"ps",
		"--filter",
		"status=exited",
		"-q",
	)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return flaw.From(err)
	}

	outputString := strings.TrimSpace(string(stdoutStderr))
	outputSlice := strings.Split(outputString, "\n")

	// remove exited containers
	if len(outputSlice) > 1 {
		for _, num := range outputSlice {
			cmd := exec.Command(
				"docker",
				"rm",
				"-f",
				num,
			)
			stdoutStderr, err = cmd.CombinedOutput()
			if err != nil {
				return flaw.From(err)
			}
		}
	}

	return nil
}
