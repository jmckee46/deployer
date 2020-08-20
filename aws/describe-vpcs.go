package awsfuncs

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jmckee46/deployer/flaw"
)

// DescribeVpcs
func DescribeVpcs(state *State) (*ec2.DescribeVpcsOutput, flaw.Flaw) {
	// get aws client
	ec2Cli := state.AWSClient.EC2

	// create input struct
	input := &ec2.DescribeVpcsInput{
		// VpcIds: []*string{
		// 	aws.String("vpc-a01106c2"),
		// },
	}

	// describe vpcs
	output, err := ec2Cli.DescribeVpcs(input)
	if err != nil {
		return nil, flaw.From(err)
	}
	// fmt.Println("\n  describe vpcs output:", output)

	return output, nil
}
