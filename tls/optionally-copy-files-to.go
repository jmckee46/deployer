package tlsDeployer

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

// OptionallyCopyFiles copies tls files to the package path if they exist and
// creates them if they don't exist
func OptionallyCopyFiles(repositoryRoot, packagePath string) flaw.Flaw {
	fmt.Printf("  optionally-copy-tls-files-to   %s...\n", packagePath)

	packageTLSPath := packagePath + "/tls/files"

	err := os.MkdirAll(packageTLSPath, 0755)
	if err != nil {
		return flaw.From(err).Wrap("tls optionally copy files failed")
	}

	dirExists, _ := osfuncs.Exists(repositoryRoot + "/tls/files")
	if dirExists {
		_, err = osfuncs.CopyFile(repositoryRoot+"/tls/files/certificate-chain.pem", packageTLSPath)
		if err != nil {
			return flaw.From(err).Wrap("tls optionally copy files failed")
		}

		_, err = osfuncs.CopyFile(repositoryRoot+"/tls/files/private-key.pem", packageTLSPath)
		if err != nil {
			return flaw.From(err).Wrap("tls optionally copy files failed")
		}
	} else {
		_, err = os.Create(packagePath + "/tls/files/certificate-chain.pem")
		if err != nil {
			return flaw.From(err).Wrap("tls optionally copy files failed")
		}
		_, err = os.Create(packagePath + "/tls/files/private-key.pem")
		if err != nil {
			return flaw.From(err).Wrap("tls optionally copy files failed")
		}
	}

	return nil
}
