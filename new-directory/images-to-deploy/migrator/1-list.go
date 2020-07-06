package main

import (
	"io/ioutil"

	"github.com/jmckee46/deployer/env-vars"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func list(state *state) {
	logger.Debug("migrator-list", state)

	migrationFiles, err := ioutil.ReadDir(envvars.MigrationsPath)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot list")
		logger.Panic("migrator-list-err", state)
	}

	state.Files = migrationFiles

	iterate(state)
}
