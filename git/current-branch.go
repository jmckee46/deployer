package git

import (
	"os/exec"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// CurrentBranch returns the current git branch
func CurrentBranch() (string, flaw.Flaw) {
	branchNameBytes, err := exec.Command(
		"git",
		"rev-parse",
		"--abbrev-ref",
		"HEAD",
	).Output()

	if err != nil {
		return "", flaw.From(err).Wrap("cannot CurrentBranch")
	}

	return strings.TrimSpace(string(branchNameBytes)), nil
}
