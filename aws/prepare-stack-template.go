package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// PrepareStackTemplate converts a template file to a string
func PrepareStackTemplate() flaw.Flaw {
	fmt.Println("Preparing stack...")

	// render stack
	err := RenderStackTemplate()
	if err != nil {
		return err
	}

	// put template in s3
	s3Dir := filepath.Join(
		os.Getenv("DE_GIT_BRANCH"),
		"templates",
		os.Getenv("DE_GIT_SHA"),
	) + ".template"

	localFile := filepath.Join(
		os.Getenv("DE_ARTIFACTS_PATH"),
		os.Getenv("DE_GIT_BRANCH"),
		"completeStack",
	)

	err = PutFileInS3(s3Dir, localFile)
	if err != nil {
		return err
	}

	// validate stack template
	err = ValidateStackTemplate()
	if err != nil {
		return err
	}

	return nil
}
