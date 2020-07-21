package tlsDeployer

import (
	"github.com/aws/aws-sdk-go/service/s3"
	awsclient "github.com/jmckee46/deployer/aws/client"
)

type state struct {
	S3Cli awsInterface
}

type awsInterface interface {
	ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
}

func newState() *state {

	state := &state{
		S3Cli: awsclient.FromPool().S3,
	}

	return state
}
