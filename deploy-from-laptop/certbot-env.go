package main

import (
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

func initCertbotEnv() flaw.Flaw {
	// set up cerbot-env
	certbotBytes, err := exec.Command(
		"/bin/sh",
		"-c",
		"virtualenv",
		"certbot-env",
		"source",
		"certbot-env/bin/activate",
	).Output()

	fmt.Println("certbotString:", string(certbotBytes))
	if err != nil {
		return flaw.From(err).Wrap("cannot set up certbot")
	}

	// // add local directory to $PATH to facilitte activation
	// err = os.Setenv("PATH", os.Getenv("PATH")+":"+os.Getenv("PWD"))
	// if err != nil {
	// 	return flaw.From(err).Wrap("cannot set up certbot")
	// }
	// fmt.Println("path:", os.Getenv("PATH"))

	// // activate certbot-env
	// certbotBytes, err = exec.Command(
	// 	"ls",
	// 	// "certbot-env/bin/activate",
	// ).Output()

	// fmt.Println("certbotString3:", string(certbotBytes))
	// if err != nil {
	// 	return flaw.From(err).Wrap("cannot set up certbot")
	// }

	return nil
}
