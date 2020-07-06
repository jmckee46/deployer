package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func execute(state *state) {
	logger.Debug("migrator-execute", state)

	output := bytes.NewBuffer([]byte{})

	cmd := exec.Cmd{
		Path:   state.FilepathName,
		Stdout: output,
		Stderr: output,
	}

	err := cmd.Run()

	if err != nil {
		outputBytes, readErr := ioutil.ReadAll(output)

		if readErr != nil {
			state.FlawError = flaw.From(readErr).Wrap("could not read migration output")
			logger.Panic("migrator-execute-failed", state)
		}

		logger.Panic("migrator-execute-failed", string(outputBytes))
	}

	record(state)
}
