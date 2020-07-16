package awsfuncs

import (
	"io/ioutil"
	"os"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

// TemplateToString converts a template file to a string
func TemplateToString(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		logger.Panic("template-to-string", flaw.From(err))
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Panic("template-to-string", flaw.From(err))
	}

	return string(b)
}
