package awsfuncs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
)

func GetTlsFilesFromS3() flaw.Flaw {
	fmt.Println("getting tls files from S3...")
	fmt.Println("")

	// get aws client
	awsS3 := awsclient.FromPool().S3

	// create uploader
	// uploader := s3manager.NewUploaderWithClient(awsS3)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Prefix: aws.String(os.Getenv("DE_GIT_BRANCH") + "/tls"),
	}

	resp, err := awsS3.ListObjectsV2(input)
	if err != nil {
		fmt.Println("err:", err)
		return flaw.From(err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	return nil
}
