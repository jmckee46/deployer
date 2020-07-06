package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

// CurrentSha returns the current git shaa
func CurrentSha() string {
	output, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		logger.Panic("CurrentSha", flaw.From(err))
	}
	outputString := fmt.Sprintf("%s", output)
	outputString = strings.TrimSpace(outputString)

	return outputString
}
