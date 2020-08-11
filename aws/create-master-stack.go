package awsfuncs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/flaw"
)

// CreateMasterStack creates the initial master stack
func CreateMasterStack(state *state) flaw.Flaw {
	// get aws client
	cloud := state.AWSClient.Cloudform

	// create input struct
	input := &cloudformation.CreateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_IAM"), aws.String("CAPABILITY_NAMED_IAM")},
		Parameters:   state.StackParametersAll,
		StackName:    aws.String(os.Getenv("DE_STACK_NAME")),
		TemplateURL:  aws.String(state.RenderedTemplateS3URL),
	}

	// create master stack
	_, err := cloud.CreateStack(input)
	if err != nil {
		return flaw.From(err)
	}
	// fmt.Println("\n  template validation output:", output)

	return nil
}
