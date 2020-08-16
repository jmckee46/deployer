package docker

import (
	"fmt"
	"os"
	"os/exec"

	awsfunc "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/flaw"
)

// Login
func Login(state *awsfunc.State) flaw.Flaw {
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
		state.GetDockerRegistry(),
	)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
		return flaw.From(err)
	}

	return nil
}
