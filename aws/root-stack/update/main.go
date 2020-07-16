package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsfuncs "github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	fmt.Println("updating root stack...")
	fmt.Println("")

	// get aws client
	cloudForm := awsclient.FromPool().Cloudform

	// create the Update Stack Input struct
	rootBucketParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeRootBucket"),
		ParameterValue: aws.String(os.Getenv("DE_ROOT_BUCKET")),
	}

	stackInput := &cloudformation.UpdateStackInput{
		Parameters:   []*cloudformation.Parameter{&rootBucketParameter},
		TemplateBody: aws.String(awsfuncs.TemplateToString("aws/root-stack/root-template.json")),
		StackName:    aws.String("ROOT"),
	}

	// create the stack
	stackOutput, err := cloudForm.UpdateStack(stackInput)
	fmt.Println("stackOutput:", stackOutput)
	if err != nil {
		fmt.Println("err:", err)
		logger.Panic("create-root-stack", err)
	}
}
