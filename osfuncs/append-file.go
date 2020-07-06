package osfuncs

import (
	"io"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// AppendFile appends the fileToCopy to the destination file
func AppendFile(dstFile string, fileToCopy string) (int64, flaw.Flaw) {
	// open file to copy
	ftc, err := os.Open(fileToCopy)
	if err != nil {
		return 0, flaw.From(err)
	}
	defer ftc.Close()

	// open destination file
	destination, err := os.OpenFile(dstFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	// append data
	nBytes, err := io.Copy(destination, ftc)
	if err != nil {
		return 0, flaw.From(err)
	}

	return nBytes, nil
}
