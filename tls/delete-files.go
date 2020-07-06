package tlsDeployer

import (
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// DeleteFiles deletes the tls directory from a given path
func DeleteFiles(path string) flaw.Flaw {

	err := os.RemoveAll(path + "/tls")
	if err != nil {
		return flaw.From(err).Wrap("tls delete failed")
	}

	return nil
}
