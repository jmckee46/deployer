package docker

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jmckee46/deployer/flaw"
)

// Login logs docker into ecr registry
func Login(dockerRegistry string) flaw.Flaw {
	fmt.Println("    logging docker into ECR...")
	// get login password
	cmd := exec.Command(
		"aws",
		"ecr",
		"get-login-password",
		"--region",
		os.Getenv("AWS_REGION"),
	)
	password, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
		return flaw.From(err)
	}

	// log docker into registry
	cmd = exec.Command(
		"docker",
		"login",
		"--username",
		"AWS",
		"--password",
		string(password),
		dockerRegistry,
	)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
		return flaw.From(err)
	}

	return nil
}
