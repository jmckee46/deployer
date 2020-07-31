package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/flaw"
)

// ValidateStackTemplate validates the stack template
func ValidateStackTemplate(state *state) flaw.Flaw {
	fmt.Println("  validating stack template...")

	// get aws client
	cloud := state.AWSClient.Cloudform

	// create the ValidateTemplateInput struct
	state.renderedTemplateS3URL = "https://s3.amazonaws.com/" + filepath.Join(
		os.Getenv("DE_ROOT_BUCKET"),
		state.renderedTemplateS3,
	)

	input := &cloudformation.ValidateTemplateInput{
		TemplateURL: aws.String(state.renderedTemplateS3URL),
	}

	// validate the template
	_, err := cloud.ValidateTemplate(input)
	if err != nil {
		return flaw.From(err)
	}
	// fmt.Println("\n  template validation output:", output)

	return nil
}
