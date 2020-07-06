package awshttpsigningclient

import (
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func authentication(state *State) {
	if !fillFromEnvKeys(state) &&
		!fillFromTaskRole(state) {
		logger.Panic(
			"awshttpsigningclient-authentication",
			flaw.New("no credentials available").Wrap("cannot authenticate"),
		)
	}

	requestTimes(state)
}
