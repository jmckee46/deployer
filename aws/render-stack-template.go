package awsfuncs

import (
	"os"

	"github.com/alecthomas/template"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// RenderStackTemplate assembles the various stack templates into one
func RenderStackTemplate() flaw.Flaw {
	branchStackTemplate := template.New("root.gotemplate")

	branchStackTemplate.Delims("{{{", "}}}")

	_, err :=
		branchStackTemplate.ParseGlob("aws/stack-templates/*.gotemplate")

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	err = branchStackTemplate.Execute(os.Stdout, nil)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return nil
}
