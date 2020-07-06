package main

import (
	"regexp"

	"github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/stamps"
)

func iterate(state *state) {
	logger.Debug("migrator-iterate", state)

	IDRegexp := regexp.MustCompile(`^\d+`)

	for _, state.File = range state.Files {
		state.CreatedAt = stamps.New().String()
		state.FileName = state.File.Name()
		state.FilepathName = envvars.MigrationsPath + "/" + state.FileName
		state.FileID = IDRegexp.FindString(state.FileName)

		filter(state)
	}
}
