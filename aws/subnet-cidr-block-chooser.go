package awsfuncs

import (
	"strconv"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// DescribeVpcs
func SubnetCidrBlockChooser(state *State) flaw.Flaw {
	cidrSlice := strings.Split(state.VpcCidrBase, ".")
	thirdByte := cidrSlice[2]
	thirdByteInt, _ := strconv.ParseInt(thirdByte, 0, 64)

	cidrBlock0 := cidrSlice[0] + "." + cidrSlice[1] + "." + thirdByte + ".0/24"

	thirdByteString := strconv.FormatInt(thirdByteInt+1, 10)
	cidrBlock1 := cidrSlice[0] + "." + cidrSlice[1] + "." + thirdByteString + ".0/24"

	state.SubnetCidrBlocks = cidrBlock0 + "," + cidrBlock1

	return nil
}

// 192.168.0.0/24,192.168.1.0/24,192.168.2.0/24,192.168.3.0/24
