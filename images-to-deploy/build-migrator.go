package imagesToDeploy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/image"
	"github.com/jmckee46/deployer/osfuncs"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/gofuncs"
)

// BuildMigrator builds the migrator
func BuildMigrator() flaw.Flaw {
	// Get current working directory
	curDir, err := os.Getwd()
	if err != nil {
		flaw.From(err)
	}

	servicePath := "images-to-deploy/migrator"
	executablePath := filepath.Join(curDir, "images-to-deploy/migrator/executables")
	migrationBase := "migrations"

	// create executable directory
	err = os.MkdirAll(executablePath, 0755)
	if err != nil {
		return flaw.From(err)
	}

	// read files in the migrations directory
	files, err := ioutil.ReadDir(migrationBase)
	if err != nil {
		return flaw.From(err)
	}

	// iterate files and build migration directories
	for _, f := range files {
		if f.IsDir() {
			// update paths
			migrationPath := filepath.Join(migrationBase, f.Name())
			dstPath := filepath.Join(executablePath, f.Name())
			migrationFile := filepath.Join(migrationPath, f.Name())

			// compile go code
			flawErr := gofuncs.Build(migrationPath)
			if flawErr != nil {
				return flawErr
			}

			// copy executable
			_, flawErr = osfuncs.CopyFile(migrationFile, dstPath)
			if flawErr != nil {
				return flawErr
			}

			// update permissions
			err = os.Chmod(dstPath, 0755)
			if err != nil {
				return flaw.From(err)
			}

			// delete original
			err = os.Remove(migrationFile)
			if err != nil {
				return flaw.From(err)
			}
		}
	}

	// build the migrator docker container
	fmt.Println("building")
	flawErr := image.BuildGo(servicePath)
	if flawErr != nil {
		return flawErr
	}

	// remove executable directory
	flawErr = osfuncs.DeleteDirAndFiles(executablePath)
	if flawErr != nil {
		return flawErr
	}

	return nil
}
