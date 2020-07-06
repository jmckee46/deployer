package buildHttpSpecs

import (
	"github.com/jmckee46/deployer/image"

	"github.com/jmckee46/deployer/flaw"
)

// Build builds the http-spec docker image
func Build() flaw.Flaw {
	// make http-specs directory
	curDir, httpDir, err := getDirectories()
	if err != nil {
		return flaw.From(err)
	}

	// assemble the specs
	err = assembleSpecs(curDir, httpDir)
	if err != nil {
		return flaw.From(err)
	}

	// image/build-with-tls images-to-deploy/http-specs
	err = image.BuildGoWithTLS("images-to-deploy/http-specs")
	if err != nil {
		// fmt.Println("build go err:", err.String())
		return flaw.From(err)
	}

	// clean up/delete specs
	err = deleteSpecs(httpDir)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
