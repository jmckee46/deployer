package main

import (
	"fmt"

	"github.com/jmckee46/deployer/aws"
	"github.com/jmckee46/deployer/aws/client"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	fmt.Println("validating root stack...")
	fmt.Println("")

	// get aws client
	awsCli := awsclient.FromPool()

	// create the Validate Stack Input struct
	stackInput := cloudformation.ValidateTemplateInput{
		TemplateBody: aws.String(awsfuncs.TemplateToString("aws/root-stack/root-template.json")),
	}

	// create the stack
	stackOutput, err := awsCli.Cloudform.ValidateTemplate(&stackInput)
	fmt.Println("stackOutput:", stackOutput)
	if err != nil {
		fmt.Println("err:", err)
		logger.Panic("validate-root-stack", err)
	}
}
