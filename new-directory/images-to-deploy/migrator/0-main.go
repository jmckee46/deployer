package main

import (
	"github.com/jmckee46/deployer/logger"
)

func main() {
	logger.Debug("migrator-main", ' ')

	state := newState()

	list(state)
}
