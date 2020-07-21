package tlsDeployer

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type mockS3 struct{}

func (m *mockS3) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	object1 := &s3.Object{
		Key:          aws.String("acm-certificate-arn"),
		LastModified: aws.Time(time.Now()),
	}
	output := &s3.ListObjectsV2Output{
		Contents: []*s3.Object{object1},
	}

	return output, nil
}

func TestTLSFiles(t *testing.T) {
	fmt.Println("calling tls test file")

	state := &state{
		S3Cli: &mockS3{},
	}

	result := tlsFilesInS3Fresh(state)

	fmt.Println("result:", result)
}
