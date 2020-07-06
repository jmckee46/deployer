package envvars

import (
	"os"

	"github.com/jmckee46/deployer/flaw"
)

func Get(envVar string) string {
	value := os.Getenv(envVar)

	if value == "" {
		flaw.New(envVar + " not set!").Panic()
	}

	return os.Getenv(envVar)
}
