package gofuncs

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// DeleteExecutable deletes the executable go code
func DeleteExecutable(path string) flaw.Flaw {
	executableName := filepath.Base(path)
	fileName := path + "/" + executableName

	cmd := exec.Command(
		"rm",
		fileName,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return flaw.From(err).Wrap("DeleteExecutable failed")
	}

	return nil

}
