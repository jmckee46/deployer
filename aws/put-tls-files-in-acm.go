package awsfuncs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

// PutTLSFilesInACM imports the tls files to aws ACM
func PutTLSFilesInACM() flaw.Flaw {
	fmt.Println("putting tls files in acm...")
	fmt.Println("")

	// get aws client
	awsAcm := awsclient.FromPool().ACM

	// create the Import Certificate Input struct
	stackInput := &acm.ImportCertificateInput{
		Certificate:      certToByte("tls/files/certificate.pem"),
		CertificateChain: certToByte("tls/files/chain.pem"),
		PrivateKey:       certToByte("tls/files/private-key.pem"),
	}

	// import the certificates
	stackOutput, err := awsAcm.ImportCertificate(stackInput)
	// fmt.Println("arn:", *stackOutput.CertificateArn)
	if err != nil {
		return flaw.From(err)
	}

	flawErr := putARNInTLS(*stackOutput.CertificateArn)
	if err != nil {
		return flawErr
	}

	return nil
}

func certToByte(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		logger.Panic("cert-to-byte", flaw.From(err))
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Panic("cert-to-byte", flaw.From(err))
	}

	return b
}

func putARNInTLS(arn string) flaw.Flaw {
	f, err := os.Create("tls/files/acm-certificate-arn")
	if err != nil {
		return flaw.From(err)
	}

	_, err = f.WriteString(arn)
	if err != nil {
		f.Close()
		return flaw.From(err)
	}

	err = f.Close()
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
