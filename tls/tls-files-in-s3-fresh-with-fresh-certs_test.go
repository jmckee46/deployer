package tlsDeployer

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var listObjectsV2Function func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)

type mockS3 struct{}

func (m *mockS3) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return listObjectsV2Function(input)
}

// testing all 5 fresh certs returned - should return true
func TestTLSFilesInS3FreshWithFreshCerts(t *testing.T) {
	state := &state{
		S3Cli: &mockS3{},
	}

	listObjectsV2Function = func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
		object1 := &s3.Object{
			Key:          aws.String("test-branch/tls/acm-certificate-arn"),
			LastModified: aws.Time(time.Now().UTC()),
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
		object5 := &s3.Object{
			Key:          aws.String("test-branch/tls/private-key.pem"),
			LastModified: aws.Time(time.Now().UTC()),
		}
		output := &s3.ListObjectsV2Output{
			Contents: []*s3.Object{
				object1,
				object2,
				object3,
				object4,
				object5,
			},
		}

		return output, nil
	}

	result := tlsFilesInS3Fresh(state)

	if result != true {
		t.Errorf("files are stale when they should be fresh, got %t instead of true", result)
	}
}
