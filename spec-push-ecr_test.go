package main

import (
	"fmt"
	"testing"

	"github.com/jmckee46/deployer/aws"
)

func TestPush(t *testing.T) {
	state := awsfuncs.NewState()

	err := awsfuncs.InitializeCidrVariables(state)
	if err != nil {
		fmt.Println("err:", err)
	}

}
