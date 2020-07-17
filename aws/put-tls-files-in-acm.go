package awsfuncs

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/logger"

	"github.com/jmckee46/deployer/flaw"
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
		PrivateKey:       certToByte("fls/files/privat-key.pem"),
	}

	// import the certificates
	stackOutput, err := awsAcm.ImportCertificate(stackInput)
	fmt.Println("stackOutput:", stackOutput)
	if err != nil {
		fmt.Println("err:", err)
		logger.Panic("create-root-stack", err)
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
