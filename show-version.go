package main

import (
	"fmt"
	"os/exec"
)

func showVersion() {
	cmd := exec.Command("docker", "ps", "-a")
	// cmd.Stdout = os.Stdout
	// cmd.Stdin = os.Stdin

	if otuput, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Output: %s\n", otuput)
	}

	cmd2 := exec.Command("pwd")

	if otuput, err := cmd2.CombinedOutput(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Output: %s\n", otuput)
	}
}
