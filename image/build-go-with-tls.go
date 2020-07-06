package image

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/gofuncs"
	"github.com/jmckee46/deployer/root-tls-certificates"
	tlsDeployer "github.com/jmckee46/deployer/tls"
)

// BuildGoWithTLS compiles go code, builds the image, installs tls files then cleans up
func BuildGoWithTLS(path string) flaw.Flaw {
	fmt.Printf("compiling and building with tls  %s...\n", path)

	// compile go code
	err := gofuncs.Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// build root tls certs
	err = rootTlsCertificates.Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// copy tls files
	wdPath, goErr := os.Getwd()
	if goErr != nil {
		return flaw.From(goErr).Wrap("BuildGoWithTLS failed")
	}

	err = tlsDeployer.OptionallyCopyFiles(wdPath, path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// build image
	err = Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// delete root tls certs
	err = rootTlsCertificates.Delete(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// delete tls files
	err = tlsDeployer.DeleteFiles(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	// delete executable
	err = gofuncs.DeleteExecutable(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGoWithTLS failed")
	}

	return nil
}
