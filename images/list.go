package images

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// List returns a slice of all docker image names that are tagged to deploy
func List() ([]string, flaw.Flaw) {

	cmd := exec.Command(
		"docker",
		"image",
		"ls",
		"--filter",
		"label=com.subledger.deployer.push",
	)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return nil, flaw.From(err)
	}

	outputString := string(stdoutStderr[:])

	var outputSlice []string
	var imageNames []string

	outputSlice = strings.Split(outputString, "\n")

	for _, line := range outputSlice {
		if strings.Contains(line, "latest") {
			lineItems := strings.Split(line, " ")
			base := filepath.Base(lineItems[0])
			imageNames = append(imageNames, base)
		}
	}

	return imageNames, nil
}
