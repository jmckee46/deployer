package gofuncs

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// GetDependencies gets go dependencies and installs gometalinter
func GetDependencies() flaw.Flaw {
	fmt.Println("getting go dependencies...")

	// might need: "git config --global http.https://gopkg.in.followRedirects true"

	cmd := exec.Command(
		"go",
		"get",
		"-u",
		"github.com/alecthomas/gometalinter",
		"github.com/aws/aws-sdk-go/aws",
		"github.com/gocarina/gocsv",
		"github.com/halorium/httprouter",
		"github.com/halorium/json-iterator",
		"github.com/lib/pq",
		"github.com/jmespath/go-jmespath",
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	// fmt.Printf("output: %s\n", out.String())
	// fmt.Printf("err: %s\n", err)
	if err != nil {
		return flaw.New(out.String())
	}

	cmd = exec.Command(
		"gometalinter",
		"--install",
	)
	cmd.Stdout = &out
	err = cmd.Run()
	// fmt.Printf("output: %s\n", out.String())
	// fmt.Printf("err: %s\n", err)
	if err != nil {
		return flaw.New(out.String())
	}

	return nil
}
