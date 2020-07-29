package awsfuncs

import (
	"os"

	"github.com/alecthomas/template"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// PrepareStackTemplate converts a template file to a string
func RenderStackTemplate() flaw.Flaw {
	branchStackTemplate := template.New("template.gotemplate")

	branchStackTemplate.Delims("{{{", "}}}")

	_, err :=
		branchStackTemplate.ParseGlob("aws/branch-stack/*.gotemplate")

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	err = branchStackTemplate.Execute(os.Stdout, nil)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return nil
}
