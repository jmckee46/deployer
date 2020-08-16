package awsfuncs

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/jmckee46/deployer/logger"
)

// MasterStackExists determines if the master stack exists
func MasterStackExists(state *State) bool {

	err := DescribeStack("master", state)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Get error details
			if awsErr.Message() == "Stack with id master does not exist" {
				return false
			}

		} else {
			logger.Panic("MasterStackExists unexpected err:", err)
		}
	}
	return true
}
