package main

import (
	"fmt"

	"github.com/jmckee46/deployer/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/logger"
)

func main() {
	// create session
	sess, err := session.NewSession()
	if err != nil {
		logger.Panic("create-root-stack", err)
	}

	// create cloudformation client
	cloudForm := cloudformation.New(sess)

	// create the Validate Stack Input struct

	stackInput := cloudformation.ValidateTemplateInput{
		TemplateBody: aws.String(awsfuncs.TemplateToString("aws/root-stack/root-template.json")),
	}

	// create the stack
	stackOutput, err := cloudForm.ValidateTemplate(&stackInput)
	fmt.Println("stackOutput:", stackOutput)
	if err != nil {
		fmt.Println("err:", err)
		logger.Panic("validate-root-stack", err)
	}
}
