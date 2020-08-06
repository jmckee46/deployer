package awsfuncs

import (
	"io/ioutil"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// GetAcmCertificateArn returns the cert arn from the cert file
func GetAcmCertificateArn() (string, flaw.Flaw) {
	// open arn file
	arnFile, err := os.Open("tls/files/acm-certificate-arn")
	if err != nil {
		return "", flaw.From(err)
	}
	defer arnFile.Close()

	// read in arn
	arnBytes, err := ioutil.ReadAll(arnFile)
	if err != nil {
		return "", flaw.From(err)
	}

	return string(arnBytes), nil
}
