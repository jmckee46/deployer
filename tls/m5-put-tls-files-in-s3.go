package tlsDeployer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
)

// PutTLSFilesInS3 uploads local TLS files to S3
func PutTLSFilesInS3() flaw.Flaw {
	fmt.Println("  putting tls files in S3...")

	// get aws client
	awsS3 := awsclient.FromPool().S3

	// create uploader
	uploader := s3manager.NewUploaderWithClient(awsS3)

	// read tls/files directory
	filesToUpload, err := readDir("tls/files/")
	if err != nil {
		return flaw.From(err)
	}

	// upload the files to s3
	for _, file := range filesToUpload {
		uploadFile(uploader, file)
	}

	return nil
}

func uploadFile(uploader *s3manager.Uploader, filename string) flaw.Flaw {
	file, err := os.Open("tls/files/" + filename)
	if err != nil {
		return flaw.From(err)
	}

	defer file.Close()

	// create the Uploader Input struct
	input := &s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String(os.Getenv("DE_ROOT_BUCKET")),
		Key:    aws.String(os.Getenv("DE_GIT_BRANCH") + "/tls/" + filename),
	}

	// upload the file
	_, err = uploader.Upload(input)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}

func readDir(dir string) ([]string, flaw.Flaw) {
	var files []string
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, flaw.From(err)
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
