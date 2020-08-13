package buildHttpSpecs

import (
	"os"

	"github.com/jmckee46/deployer/flaw"
)

func getDirectories() (string, string, flaw.Flaw) {
	// Get current working directory
	curDir, err := os.Getwd()
	if err != nil {
		return "", "", flaw.From(err)
	}

	// get http-specs directory
	httpDir := curDir + "/images-to-deploy/http-specs"

	return curDir, httpDir, nil
}
