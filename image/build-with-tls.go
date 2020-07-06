package image

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/root-tls-certificates"
	tlsDeployer "github.com/jmckee46/deployer/tls"
)

// BuildWithTLS performs a docker build on the path passed and installs tls files
func BuildWithTLS(path string) flaw.Flaw {
	fmt.Printf("building with tls                %s...\n", path)

	// build root tls certs
	err := rootTlsCertificates.Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	// copy tls files
	wdPath, goErr := os.Getwd()
	if goErr != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	err = tlsDeployer.OptionallyCopyFiles(wdPath, path)
	if err != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	// build image
	err = Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	// delete root tls certs
	err = rootTlsCertificates.Delete(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	// delete tls files
	err = tlsDeployer.DeleteFiles(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildWithTLS failed")
	}

	return nil
}
