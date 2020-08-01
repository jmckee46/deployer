package awsfuncs

import (
	"os"
	"strings"

	"github.com/jmckee46/deployer/git"
	"github.com/jmckee46/deployer/logger"
)

// NotOKToDeploy checks if stack name and head commit allows for deploying or not.
// If on master it should always deploy regardless of commit, and function should return false.
// Otherwise, non master branches need to end in [deploy] to deploy, and function should return false.
func NotOKToDeploy() bool {
	headCommitMsg, err := git.HeadCommit()
	if err != nil {
		logger.Panic("CheckOKToDeploy", err)
	}

	if os.Getenv("DE_STACK_NAME") != "master" &&
		!strings.HasSuffix(headCommitMsg, "[deploy]") {
		return true
	}

	return false
}
