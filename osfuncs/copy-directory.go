package osfuncs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// CopyDirectory copies a directory including all subdirectories
// Example, src = foo/bar, dst = newDir. Result newDir/bar.
// If bar is a directory, copyDirectory will create the directory automatically.
func CopyDirectory(src, dst string) flaw.Flaw {
	newDirBase := filepath.Base(src)
	newDir := filepath.Join(dst, newDirBase)

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return flaw.From(err)
	}

	err = os.MkdirAll(newDir, 0755)
	if err != nil {
		return flaw.From(err)
	}

	if len(files) > 0 {
		for _, file := range files {
			newSrc := filepath.Join(src, file.Name())
			newDst := filepath.Join(newDir, file.Name())
			if file.IsDir() {
				CopyDirectory(newSrc, newDir)
				continue
			}
			_, copyErr := CopyFile(newSrc, newDst)
			if copyErr != nil {
				return flaw.From(copyErr)
			}
		}
	}

	return nil
}
