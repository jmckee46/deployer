package osfuncs

import (
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
)

// DeleteDirAndFiles deletes all files in a directory then deletes the directory
func DeleteDirAndFiles(dir string) flaw.Flaw {
	d, err := os.Open(dir)
	if err != nil {
		return flaw.From(err)
	}

	names, err := d.Readdirnames(-1)
	if err != nil {
		return flaw.From(err)
	}

	d.Close()

	// remove files in directory
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return flaw.From(err)
		}
	}

	// remove directory
	err = os.RemoveAll(dir)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
