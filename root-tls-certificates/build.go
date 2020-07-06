package rootTlsCertificates

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

// Build copies the root tls certs and appends the mock load-balancer to it
func Build(path string) flaw.Flaw {
	// get the list of trusted certs
	destPath := path + "/ca-certificates.crt"
	if runtime.GOOS == "darwin" {
		cmd := exec.Command(
			"cp",
			"/usr/local/etc/openssl/cert.pem",
			destPath,
		)
		_, err := cmd.CombinedOutput()

		if err != nil {
			return flaw.From(err).Wrap("root-tls-certs build failed")
		}
	} else {
		cmd := exec.Command(
			"cp",
			"/etc/ssl/certs/ca-certificates.crt",
			destPath,
		)
		_, err := cmd.CombinedOutput()

		if err != nil {
			return flaw.From(err).Wrap("root-tls-certs build failed")
		}
	}

	// Get current working directory
	curDir, err := os.Getwd()
	if err != nil {
		return flaw.From(err)
	}

	//  append the mock load-balancer cert to the list of trusted certs
	ftc := filepath.Join(curDir, "/docker-compose-only-images/load-balancer/load-balancer.crt")
	_, flawErr := osfuncs.AppendFile(destPath, ftc)
	if flawErr != nil {
		return flawErr
	}

	return nil
}
