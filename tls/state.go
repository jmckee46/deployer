package tlsDeployer

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	awsclient "github.com/jmckee46/deployer/aws/client"
)

type state struct {
	S3Cli s3iface.S3API
}

// type awsInterface interface {
// 	ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
// 	NewDownloaderWithClient(svc s3iface.S3API, options ...func(*s3iface.Downloader)) *s3iface.Downloader
// }

func newState() *state {

	state := &state{
		S3Cli: awsclient.FromPool().S3,
	}

	return state
}
