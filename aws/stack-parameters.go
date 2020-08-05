package awsfuncs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/flaw"
)

// StackParameters returns the stack parameters required to create/update a stack
func StackParameters(state *state) ([]*cloudformation.Parameter, flaw.Flaw) {
	fmt.Println("gathering stack parameters...")

	return nil, nil
}
