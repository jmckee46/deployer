package tlsDeployer

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/logger"
)

func tlsFilesInS3Fresh() bool {
	fmt.Println("checking tls files in S3 are fresh...")
	fmt.Println("")

	// get aws client
	awsS3 := awsclient.FromPool().S3

	prefix := os.Getenv("DE_GIT_BRANCH") + "/tls"

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Prefix: aws.String(prefix),
	}

	resp, err := awsS3.ListObjectsV2(input)
	if err != nil {
		logger.Panic("TlsFilesInS3Fresh", err)
	}

	var acmCertificateArn, certificateChain, certificate, chain, privateKey time.Time

	if len(resp.Contents) < 1 {
		fmt.Println("tls files are not in s3...")
		return false
	}

	for _, item := range resp.Contents {
		switch *item.Key {
		case prefix + "acm-certificate-arn":
			acmCertificateArn = *item.LastModified
		case prefix + "certificate-chain.pem":
			certificateChain = *item.LastModified
		case prefix + "certificate.pem":
			certificate = *item.LastModified
		case prefix + "chain.pem":
			chain = *item.LastModified
		case prefix + "private - key.pem":
			privateKey = *item.LastModified
		}
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	if acmCertificateArn.After(sevenDaysAgo) ||
		certificateChain.After(sevenDaysAgo) ||
		certificate.After(sevenDaysAgo) ||
		chain.After(sevenDaysAgo) ||
		privateKey.After(sevenDaysAgo) {

		fmt.Println("tls files in s3 are stale...")
		return false
	}

	fmt.Println("tls files in s3 are fresh...")

	return true
}
