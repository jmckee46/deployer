package loadBalancer

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// CreateCert creates new load balancer certs
func CreateCert() flaw.Flaw {
	fmt.Println("creating load-balancer cert...")

	cmd := exec.Command("openssl",
		"req",
		"-x509",
		"-out",
		"docker-compose-only-images/load-balancer/load-balancer.crt",
		"-keyout",
		"docker-compose-only-images/load-balancer/load-balancer.key",
		"-newkey",
		"rsa:2048",
		"-nodes",
		"-sha256",
		"-subj",
		"/CN=load-balancer",
		"-extensions",
		"EXT",
		"-config",
		"docker-compose-only-images/load-balancer/config-file",
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if outStr != "" {
		fmt.Printf("out:\n%s\n", outStr)
	}
	if errStr != "" {
		fmt.Printf("%s\n", errStr)
	}
	if err != nil {
		return flaw.From(err)
		// log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return nil
}
