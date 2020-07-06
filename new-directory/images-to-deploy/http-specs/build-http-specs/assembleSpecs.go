package buildHttpSpecs

import (
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

func assembleSpecs(curDir string, httpDir string) flaw.Flaw {
	// gather .htsf files
	var files []string

	err := filepath.Walk(curDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".htsf" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return flaw.From(err)
	}

	// copy .htsf files to http-specs
	for _, file := range files {
		// compute destination directory
		shortPath, _ := filepath.Rel(curDir, file)
		dst := filepath.Join(httpDir, shortPath)
		dstDir := filepath.Dir(dst)

		// create destination directory
		err = os.MkdirAll(dstDir, 0755)
		if err != nil {
			return flaw.From(err)
		}

		// copy the file
		_, err := osfuncs.CopyFile(file, dst)
		if err != nil {
			return flaw.From(err)
		}
	}

	return nil
}
