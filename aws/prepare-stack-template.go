package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// PrepareStackTemplate converts a template file to a string
func PrepareStackTemplate(state *State) flaw.Flaw {
	fmt.Println("Preparing stack template...")

	// render stack
	err := RenderStackTemplate(state)
	if err != nil {
		return err
	}

	// validate target group names
	err = ValidateTargetGroupNames(state)
	if err != nil {
		return err
	}

	// set deletion policy
	err = SetDeletionPolicy(state)
	if err != nil {
		return err
	}

	// put template in s3
	state.RenderedTemplateS3 = filepath.Join(
		os.Getenv("DE_GIT_BRANCH"),
		"templates",
		os.Getenv("DE_GIT_SHA"),
	) + ".template"

	err = PutFileInS3(state, state.RenderedTemplateS3, state.RenderedTemplateLocal)
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
