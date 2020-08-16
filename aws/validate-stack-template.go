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
func ValidateStackTemplate(state *State) flaw.Flaw {
	fmt.Println("  validating stack template...")

	// get aws client
	cloud := state.AWSClient.Cloudform

	// create the ValidateTemplateInput struct
	state.RenderedTemplateS3URL = "https://s3.amazonaws.com/" + filepath.Join(
		os.Getenv("DE_ROOT_BUCKET"),
		state.RenderedTemplateS3,
	)

	input := &cloudformation.ValidateTemplateInput{
		TemplateURL: aws.String(state.RenderedTemplateS3URL),
	}

	// validate the template
	_, err := cloud.ValidateTemplate(input)
	if err != nil {
		return flaw.From(err)
	}
	// fmt.Println("\n  template validation output:", output)

	return nil
}
