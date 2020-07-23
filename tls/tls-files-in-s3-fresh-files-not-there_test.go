package tlsDeployer

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

// testing no files returned - should return false
func TestTLSFilesInS3FreshWithFilesNotThere(t *testing.T) {
	state := &state{
		S3Cli: &mockS3{},
	}

	listObjectsV2Function = func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
		output := &s3.ListObjectsV2Output{}

		return output, nil
	}

	result := tlsFilesInS3Fresh(state)

	if result != false {
		t.Errorf("files are fresh when they should be stale, got %t instead of false", result)
	}
}
