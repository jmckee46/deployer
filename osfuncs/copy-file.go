package osfuncs

import (
	"fmt"
	"io"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// CopyFile copies a file
func CopyFile(src, dst string) (int64, flaw.Flaw) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, flaw.From(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		info := fmt.Sprintf("%s is not a regular file", src)
		return 0, flaw.New(info)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, flaw.From(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, flaw.From(err)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		return 0, flaw.From(err)
	}

	return nBytes, nil
}
