package certbot

import (
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/osfuncs"

	"github.com/jmckee46/deployer/flaw"
)

// CopyTLSFilesToTLSDirectory copies certbot certs to tls/files directory
func CopyTLSFilesToTLSDirectory() flaw.Flaw {
	// create tls/files directory
	err := os.MkdirAll("tls/files", 0755)
	if err != nil {
		return flaw.From(err)
	}

	// copy cert.pem
	certbotBase := filepath.Join("certbot/files/live", os.Getenv("DE_DOMAIN"))
	certbotCertPem := filepath.Join(certbotBase, "cert.pem")
	_, flawErr := osfuncs.CopyFile(certbotCertPem, "tls/files/certificate.pem")
	if flawErr != nil {
		return flawErr
	}

	// copy chain.pem
	certbotChainPem := filepath.Join(certbotBase, "chain.pem")
	_, flawErr = osfuncs.CopyFile(certbotChainPem, "tls/files/chain.pem")
	if flawErr != nil {
		return flawErr
	}

	// copy privkey.pem
	certbotPrivkeyPem := filepath.Join(certbotBase, "privkey.pem")
	_, flawErr = osfuncs.CopyFile(certbotPrivkeyPem, "tls/files/private-key.pem")
	if flawErr != nil {
		return flawErr
	}

	return nil
}
