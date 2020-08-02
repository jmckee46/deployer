package awsfuncs

import (
	awsclient "github.com/jmckee46/deployer/aws/client"
)

type state struct {
	AWSClient             *awsclient.Client
	RenderedTemplateLocal string
	RenderedTemplateS3    string
	RenderedTemplateS3URL string
}

func NewState() *state {

	state := &state{
		AWSClient: awsclient.FromPool(),
	}

	return state
}
