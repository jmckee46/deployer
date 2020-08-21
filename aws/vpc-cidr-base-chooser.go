package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// VpcCidrBaseChooser initializes the vpcCidrBase
func VpcCidrBaseChooser(state *State) flaw.Flaw {
	output, err := DescribeVpcs(state)
	if err != nil {
		return err
	}
	var cidrBlocksInUse []string

	for _, vpc := range output.Vpcs {
		cidrBlocksInUse = append(cidrBlocksInUse, *vpc.CidrBlock)
	}

	sequence := []string{"0", "4", "252"}

	for _, thirdByte := range sequence {
		proposedCidrBase := "192.168." + thirdByte

		for _, cidrBlockInUse := range cidrBlocksInUse {
			if proposedCidrBase == cidrBlockInUse {
				fmt.Println(proposedCidrBase, "is already in use...")
			} else {
				state.VpcCidrBase = proposedCidrBase
				return nil
			}
		}
	}

	return flaw.New("ran out of acceptable cidr bases")
}
