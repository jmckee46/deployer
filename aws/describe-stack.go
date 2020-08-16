package awsfuncs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// DescribeStack
func DescribeStack(stackName string, state *State) error {
	// get aws client
	cloud := state.AWSClient.Cloudform

	// create input struct
	input := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}

	// describe stack
	_, err := cloud.DescribeStacks(input)
	if err != nil {
		return err
	}
	// fmt.Println("\n  describe stack output:", output)

	return nil
}
