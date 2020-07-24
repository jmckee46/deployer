package tlsDeployer

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jmckee46/deployer/flaw"
)

func getTLSFilesFromS3(state *state) flaw.Flaw {
	fmt.Println("getting tls files from S3...")
	fmt.Println("")

	downloader := s3manager.NewDownloaderWithClient(state.S3Cli)

	filesToDownload := []string{
		"acm-certificate-arn",
		"certificate-chain.pem",
		"certificate.pem",
		"chain.pem",
		"private-key.pem",
	}

	s3Prefix := os.Getenv("DE_GIT_BRANCH") + "/tls/"
	localPrefix := "tls/files/"

	for _, filename := range filesToDownload {
		s3FileName := s3Prefix + filename
		localFilename := localPrefix + filename

		err := getFile(downloader, s3FileName, localFilename)
		if err != nil {
			return err
		}

	}

	return nil
}

func getFile(downloader *s3manager.Downloader, s3FileName string, localFilename string) flaw.Flaw {
	// create file to write to
	f, err := os.Create(localFilename)
	if err != nil {
		flaw.From(err)
	}

	// write contents of s3 object to file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Key:    aws.String(s3FileName),
	})
	if err != nil {
		flaw.From(err)
	}

	f.Close()

	return nil
}
