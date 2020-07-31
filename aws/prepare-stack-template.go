package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// PrepareStackTemplate converts a template file to a string
func PrepareStackTemplate(state *state) flaw.Flaw {
	fmt.Println("Preparing stack...")

	// render stack
	err := RenderStackTemplate(state)
	if err != nil {
		return err
	}

	// put template in s3
	state.renderedTemplateS3 = filepath.Join(
		os.Getenv("DE_GIT_BRANCH"),
		"templates",
		os.Getenv("DE_GIT_SHA"),
	) + ".template"

	err = PutFileInS3(state, state.renderedTemplateS3, state.renderedTemplateLocal)
	if err != nil {
		return err
	}

	// validate stack template
	err = ValidateStackTemplate(state)
	if err != nil {
		return err
	}

	return nil
}
