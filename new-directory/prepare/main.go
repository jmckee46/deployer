package main

import (
	"fmt"

	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/myAppTest/docker"
	"github.com/jmckee46/myAppTest/docker-compose"
	"github.com/mgutz/ansi"
)

func main() {
	Red := ansi.ColorCode("red")
	Green := ansi.ColorCode("green")
	Reset := ansi.ColorCode("reset")

	err := docker.Build()
	if err != nil {
		fmt.Println(string(Red), "PREPARE FAILED!", string(Reset), "🤔")
		logger.Panic("myAppTest build err", err.String())
	}

	err = dockerCompose.Start()
	if err != nil {
		fmt.Println(string(Red), "PREPARE FAILED!", string(Reset), "🤔")
		logger.Panic("myAppTest start err", err.String())
	}

	err = dockerCompose.Test()
	if err != nil {
		fmt.Println(string(Red), "PREPARE FAILED!", string(Reset), "🤔")
		logger.Panic("myAppTest test err", err.String())
	}

	fmt.Println("")
	fmt.Println(string(Green), "PREPARE SUCCEEDED!", string(Reset), "😎")

}
