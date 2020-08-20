package awsfuncs

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	awsclient "github.com/jmckee46/deployer/aws/client"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

type State struct {
	AWSAccountID              string
	AWSClient                 *awsclient.Client
	DockerRegistry            string
	ImagesToDeploy            []string
	RenderedTemplateLocal     string
	RenderedTemplateS3        string
	RenderedTemplateS3URL     string
	StackParametersAll        []*cloudformation.Parameter
	StackParametersStackState []*cloudformation.Parameter
	StackParametersPasswords  []*cloudformation.Parameter
	SubnetCidrBlocks          string
	VpcCidrBase               string
}

func NewState() *State {

	state := &State{
		AWSClient: awsclient.FromPool(),
	}

	return state
}

func (s *State) GetAWSAccountID() string {
	if s.AWSAccountID == "" {
		// get account info
		cmd := exec.Command(
			"aws",
			"sts",
			"get-caller-identity",
		)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			logger.Panic("state.GetAWSAccountID()", flaw.From(err))
		}

		// unmarshal the json
		type awsIdentity struct {
			UserID  string
			Account string
			Arn     string
		}
		result := &awsIdentity{}
		err = json.Unmarshal(stdoutStderr, &result)
		s.AWSAccountID = result.Account

		return result.Account
	}

	return s.AWSAccountID
}

func (s *State) GetDockerRegistry() string {
	if s.DockerRegistry == "" {
		dr := s.GetAWSAccountID() + ".dkr.ecr." + os.Getenv("AWS_REGION") + ".amazonaws.com"
		s.DockerRegistry = dr

		return dr
	}

	return s.DockerRegistry
}
