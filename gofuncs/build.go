package gofuncs

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Build runs go build on path provided
// TODO: The below is close, but could be better. Two issues, one exec.Command and/or "go build" seems
// to have a problem setting the env vars using cmd.Env, but works fine passing in with go build
// The second issue was -ldflags. The -s -w need to be passed together, but exec.Command parses
// by spaces and won't allow this. This problem is supposed to be addressed in a future release.
func Build(path string) flaw.Flaw {
	fmt.Printf("  compiling                      %s...\n", path)

	// docker build the go code
	// envString := []string{"GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0"}
	cmd := exec.Command(
		"env",
		"GOOS=linux",
		"GOARCH=amd64",
		"CGO_ENABLED=0",
		"go",
		"build",
		"-i",
		"-ldflags=-w",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = path
	// cmd.Env = envString
	err := cmd.Run()
	// fmt.Printf("output: %s\n", out.String())
	if err != nil {
		// fmt.Printf("err: %s\n", err)
		return flaw.From(err)
	}

	return nil
}
