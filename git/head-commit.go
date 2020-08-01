package git

import (
	"os/exec"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// HeadCommit returns the current git head commit
func HeadCommit() (string, flaw.Flaw) {
	branchNameBytes, err := exec.Command(
		"git",
		"log",
		"--oneline",
		"-1",
	).Output()

	if err != nil {
		return "", flaw.From(err).Wrap("cannot HeadCommit")
	}

	return strings.TrimSpace(string(branchNameBytes)), nil
}
