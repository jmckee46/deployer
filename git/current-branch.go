package git

import (
	"os/exec"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

func CurrentBranch() (string, flaw.Flaw) {
	branchNameBytes, err := exec.Command("git/current-branch").Output()

	if err != nil {
		return "", flaw.From(err).Wrap("cannot CurrentBranch")
	}

	return strings.TrimSpace(string(branchNameBytes)), nil
}
