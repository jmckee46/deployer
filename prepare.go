package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func prepare() {
	// Get current working directory
	curDir, err := os.Getwd()
	if err != nil {
		logger.Panic("prepare", flaw.From(err))
	}
	curDir = curDir + "/"

	// check for prepare/main.go
	filePath := curDir + "prepare/main.go"
	if fileExists(filePath) {
		cmd := exec.Command("go", "run", filePath)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
	} else {
		fmt.Println(filePath, "does not exist (or is a directory)")
	}
}
