package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
)

// ValidateStackTemplate validates the stack template
func ValidateStackTemplate() flaw.Flaw {
	fmt.Println("  validating stack template...")

	// get aws client
	cloud := awsclient.FromPool().Cloudform

	// create the ValidateTemplateInput struct
	s3Url := "https://s3.amazonaws.com/" + filepath.Join(
		os.Getenv("DE_ROOT_BUCKET"),
		os.Getenv("DE_GIT_BRANCH"),
		"templates",
		os.Getenv("DE_GIT_SHA"),
	) + ".template"

	input := &cloudformation.ValidateTemplateInput{
		TemplateURL: aws.String(s3Url),
	}

	// validate the template
	output, err := cloud.ValidateTemplate(input)
	if err != nil {
		return flaw.From(err)
	}
	if output != nil {
		fmt.Println("\n  template validation failed:", output)
	}

	return nil
}
