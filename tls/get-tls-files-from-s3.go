package tlsDeployer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jmckee46/deployer/flaw"
)

func getTLSFilesFromS3(state *state) flaw.Flaw {
	fmt.Println("getting tls files from S3...")
	fmt.Println("")

	// create uploader
	downloader := s3manager.NewDownloaderWithClient(state.S3Cli)

	// read tls/files directory
	filesToUpload, err := readDir("tls/files/")
	if err != nil {
		fmt.Println("dir err:", err)
		return flaw.From(err)
	}

	// upload the files to s3
	for _, file := range filesToUpload {
		uploadFile(uploader, file)
	}

	return nil
}
