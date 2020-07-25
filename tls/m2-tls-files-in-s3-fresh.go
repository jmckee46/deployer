package tlsDeployer

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmckee46/deployer/logger"
)

// checks all files are present and all are fresh - should return true
func tlsFilesInS3Fresh(state *state) bool {
	fmt.Println("  checking tls files in S3 are fresh...")

	// get aws client
	awsS3 := state.S3Cli

	prefix := os.Getenv("DE_GIT_BRANCH") + "/tls/"

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Prefix: aws.String(prefix),
	}

	resp, err := awsS3.ListObjectsV2(input)
	if err != nil {
		logger.Panic("TlsFilesInS3Fresh", err)
	}

	var acmCertificateArn, certificateChain, certificate, chain, privateKey time.Time

	if resp.Contents == nil {
		fmt.Println("  tls files are not in s3...")
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
		case prefix + "private-key.pem":
			privateKey = *item.LastModified
		}
	}

	sevenDaysAgo := time.Now().UTC().AddDate(0, 0, -7)

	if acmCertificateArn.After(sevenDaysAgo) &&
		certificateChain.After(sevenDaysAgo) &&
		certificate.After(sevenDaysAgo) &&
		chain.After(sevenDaysAgo) &&
		privateKey.After(sevenDaysAgo) {

		fmt.Println("  tls files in s3 are fresh...")
		return true
	}

	fmt.Println("  tls files in s3 are stale or incomplete...")

	return false
}
