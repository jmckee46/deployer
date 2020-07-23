package tlsDeployer

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

func getTLSFilesFromS3(state *state) flaw.Flaw {
	fmt.Println("getting tls files from S3...")
	fmt.Println("")

	// create uploader
	// downloader := s3manager.NewDownloaderWithClient(state.S3Cli)

	// create file to write to

	// write contents of s3 object to file

	return nil
}
