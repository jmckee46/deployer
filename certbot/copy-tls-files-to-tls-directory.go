package certbot

import (
	"fmt"
	"io"
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

	// create certificate-chain.pem from cert.pem and chain.pem
	flawErr = createCertificateChain(certbotCertPem, certbotChainPem)
	if flawErr != nil {
		return flawErr
	}

	return nil
}

func createCertificateChain(certbotCertPem string, certbotChainPem string) flaw.Flaw {
	// make sure source files exist and are regular
	certbotCertPemStat, err := os.Stat(certbotCertPem)
	if err != nil {
		return flaw.From(err)
	}

	if !certbotCertPemStat.Mode().IsRegular() {
		info := fmt.Sprintf("%s is not a regular file", certbotCertPem)
		return flaw.New(info)
	}

	certbotChainPemStat, err := os.Stat(certbotChainPem)
	if err != nil {
		return flaw.From(err)
	}

	if !certbotChainPemStat.Mode().IsRegular() {
		info := fmt.Sprintf("%s is not a regular file", certbotChainPem)
		return flaw.New(info)
	}

	// copy both files to tls/files/certificat-chain.pem
	destination, err := os.Create("tls/files/certificate-chain.pem")
	if err != nil {
		return flaw.From(err)
	}
	defer destination.Close()

	certPem, err := os.Open(certbotCertPem)
	if err != nil {
		return flaw.From(err)
	}
	defer certPem.Close()

	_, err = io.Copy(destination, certPem)
	if err != nil {
		return flaw.From(err)
	}

	chainPem, err := os.Open(certbotChainPem)
	if err != nil {
		return flaw.From(err)
	}
	defer chainPem.Close()

	_, err = io.Copy(destination, chainPem)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
