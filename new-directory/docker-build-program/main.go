package main

import (
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/myAppTest/docker"
)

func main() {
	err := docker.Build()
	if err != nil {
		logger.Panic("myAppTest build err", err.String())
	}
}
