package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

func runSpecs() flaw.Flaw {
	filesToRun := []string{
		"images-to-deploy/health-check/spec-get-health-check.htsf",
		"images-to-deploy/health-check/spec-get-health-check.htsf",
	}

	var out bytes.Buffer

	for _, file := range filesToRun {
		cmd := exec.Command(
			"http-spec",
			"-scheme",
			"https",
			"-hostname",
			"load-balancer",
			"-max-http-attempts",
			"1",
			file,
		)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Printf("run http-specs output: %s\n", out.String())

			// return flaw.From(err).Wrap("image build failed").Wrap(out.String())
			return flaw.New(out.String())
		}
	}
	fmt.Printf("%s\n", out.String())

	return nil
}
