package buildHttpSpecs

import (
	"io/ioutil"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

// Build builds the http-spec docker image
func deleteSpecs(httpDir string) flaw.Flaw {
	// read files in the http-specs directory
	files, err := ioutil.ReadDir(httpDir)
	if err != nil {
		return flaw.From(err)
	}

	// iterate files and delete proper directories
	for _, f := range files {
		if f.IsDir() && f.Name() != "build-http-specs" {
			fullPath := filepath.Join(httpDir, f.Name())
			err = osfuncs.DeleteDirAndFiles(fullPath)
			if err != nil {
				return flaw.From(err)
			}
		}
	}

	return nil
}
