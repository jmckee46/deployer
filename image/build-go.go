package image

import (
	"fmt"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/gofuncs"
)

// BuildGo compiles go code, builds the image then deletes the executable go code
func BuildGo(path string) flaw.Flaw {
	fmt.Printf("compiling and building           %s...\n", path)

	// compile go code
	err := gofuncs.Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGo failed")
	}

	// build image
	err = Build(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGo failed")
	}

	// delete executable
	err = gofuncs.DeleteExecutable(path)
	if err != nil {
		return flaw.From(err).Wrap("BuildGo failed")
	}

	return nil
}
