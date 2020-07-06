package gofuncs

import (
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

func SimpleCmd(path string, arg ...string) flaw.Flaw {
	cmd := exec.Command(
		path,
		arg...,
	)
	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("stdoutStderr: %s\n", stdoutStderr)
	if err != nil {
		// fmt.Printf("err: %s\n", err)
		return flaw.From(err)
	}

	return nil
}
