package images

import (
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Prune prunes all images not referenced by any container including dangling images
func Prune() flaw.Flaw {
	fmt.Println("pruning images...")

	cmd := exec.Command(
		"docker",
		"image",
		"prune",
		"-a",
		"-f",
	)
	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("pruning err: %s\n", stdoutStderr)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
