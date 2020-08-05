package awsfuncs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/flaw"
)

// RenderStackTemplate assembles the various stack templates into one
func CreateStackFromMaster(state *state) flaw.Flaw {
	// get aws client
	cloud := state.AWSClient.Cloudform

	// sample parameter, don't use!!!!!!!!!!!
	rootBucketParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeRootBucket"),
		ParameterValue: aws.String(os.Getenv("DE_ROOT_BUCKET")),
	}

	input := &cloudformation.CreateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_IAM"), aws.String("CAPABILITY_NAMED_IAM")},
		Parameters:   []*cloudformation.Parameter{&rootBucketParameter},
		StackName:    aws.String(os.Getenv("AN_STACK_NAME")),
		TemplateURL:  aws.String(state.RenderedTemplateS3URL),
	}

	// validate the template
	_, err := cloud.CreateStack(input)
	if err != nil {
		return flaw.From(err)
	}
	// fmt.Println("\n  template validation output:", output)

	return nil
}

// aws                                                                                                \
// cloudformation                                                                                   \
// 	create-stack                                                                                   \
// 		--capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM                                           \
// 		--stack-name $AN_STACK_NAME                                                                  \
// 		--template-url https://s3.amazonaws.com/$AN_GLOBAL_BUCKET/master/$AN_GIT_MASTER_SHA.template \
// 		--parameters "$(aws/get-master-parameters)" 2>&1
