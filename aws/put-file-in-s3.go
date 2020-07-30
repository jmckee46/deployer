package awsfuncs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
)

// PutFileInS3 uploads a single file to S3
func PutFileInS3(filename string) flaw.Flaw {
	fmt.Println("  putting file in S3...")

	// get aws client
	awsS3 := awsclient.FromPool().S3

	// create uploader
	uploader := s3manager.NewUploaderWithClient(awsS3)

	// open file
	file, err := os.Open(filename)
	if err != nil {
		return flaw.From(err)
	}

	defer file.Close()

	// create the Uploader Input struct
	input := &s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Key:    aws.String(filename),
	}

	// upload the file
	_, err = uploader.Upload(input)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
