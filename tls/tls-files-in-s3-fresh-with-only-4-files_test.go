package tlsDeployer

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// testing the returned files have a missing file, 4 fresh files - should return false
func TestTLSFilesInS3FreshWithOnly4Files(t *testing.T) {
	state := &state{
		S3Cli: &mockS3{},
	}

	listObjectsV2Function = func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
		object1 := &s3.Object{
			Key:          aws.String("test-branch/tls/acm-certificate-arn"),
			LastModified: aws.Time(time.Now().UTC().AddDate(0, 0, -14)),
		}
		object2 := &s3.Object{
			Key:          aws.String("test-branch/tls/certificate-chain.pem"),
			LastModified: aws.Time(time.Now().UTC()),
		}
		object3 := &s3.Object{
			Key:          aws.String("test-branch/tls/certificate.pem"),
			LastModified: aws.Time(time.Now().UTC()),
		}
		object4 := &s3.Object{
			Key:          aws.String("test-branch/tls/chain.pem"),
			LastModified: aws.Time(time.Now().UTC()),
		}
		output := &s3.ListObjectsV2Output{
			Contents: []*s3.Object{
				object1,
				object2,
				object3,
				object4,
			},
		}

		return output, nil
	}

	result := tlsFilesInS3Fresh(state)

	if result != false {
		t.Errorf("files are fresh when they should be stale, got %t instead of false", result)
	}
}
