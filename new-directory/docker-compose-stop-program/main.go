package main

import (
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/myAppTest/docker-compose"
)

func main() {
	err := dockerCompose.Stop()
	if err != nil {
		logger.Panic("myAppTest stop err", err.String())
	}
}
