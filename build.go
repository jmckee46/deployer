package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func build() {
	fmt.Println("deployer is building...")

	// Get current working directory
	curDir, err := os.Getwd()
	if err != nil {
		logger.Panic("initNewDirectory", flaw.From(err))
	}

	// check for build/build
	filePath := filepath.Join(curDir, "docker-build-program/main.go")

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
	} else {
		fmt.Println("no build")
		fmt.Println(filePath, "does not exist (or is a directory)")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
