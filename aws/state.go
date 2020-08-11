package awsfuncs

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsclient "github.com/jmckee46/deployer/aws/client"
)

type state struct {
	AWSClient                 *awsclient.Client
	RenderedTemplateLocal     string
	RenderedTemplateS3        string
	RenderedTemplateS3URL     string
	StackParametersAll        []*cloudformation.Parameter
	StackParametersStackState []*cloudformation.Parameter
	StackParametersPasswords  []*cloudformation.Parameter
}

func NewState() *state {

	state := &state{
		AWSClient: awsclient.FromPool(),
	}

	return state
}
