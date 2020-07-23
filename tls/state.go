package tlsDeployer

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	awsclient "github.com/jmckee46/deployer/aws/client"
)

type state struct {
	S3Cli s3iface.S3API
}

func newState() *state {

	state := &state{
		S3Cli: awsclient.FromPool().S3,
	}

	return state
}
