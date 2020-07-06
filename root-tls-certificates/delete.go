package rootTlsCertificates

import (
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// Delete deletes the root tls cert
func Delete(path string) flaw.Flaw {

	err := os.Remove(path + "/ca-certificates.crt")
	if err != nil {
		return flaw.From(err).Wrap("root-tls-certs delete failed")
	}

	return nil
}
