package awsfuncs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/template"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// RenderStackTemplate assembles the various stack templates into one
func RenderStackTemplate(state *state) flaw.Flaw {
	fmt.Println("  rendering stack template...")
	// make new branch directory in artifacts
	branchName := filepath.Join(
		os.Getenv("DE_ARTIFACTS_PATH"),
		os.Getenv("DE_GIT_BRANCH"),
	)

	err := os.MkdirAll(branchName, 0755)
	if err != nil {
		return flaw.From(err)
	}

	// create file
	state.RenderedTemplateLocal = filepath.Join(
		branchName,
		"completeStack",
	)
	newStackFile, err := os.Create(state.RenderedTemplateLocal)
	if err != nil {
		fmt.Println("trouble creating file:", err)
	}

	branchStackTemplate := template.New("root.gotemplate")

	branchStackTemplate.Delims("{{{", "}}}")

	_, err = branchStackTemplate.ParseGlob("aws/stack-templates/*.gotemplate")

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	err = branchStackTemplate.Execute(newStackFile, nil)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return nil
}
