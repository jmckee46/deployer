package awsfuncs

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
)

// InitializeCidrVariables chooses the appropriate vpc cidr base and subnet cidr blocks
func InitializeCidrVariables(state *State) flaw.Flaw {
	// must choose cidr base before choosing cidr blocks
	err := VpcCidrBaseChooser(state)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = SubnetCidrBlockChooser(state)
	if err != nil {
		fmt.Println("err:", err)
	}

	return nil
}
