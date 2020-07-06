package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/osfuncs"
)

func initNewDirectory() {
	fmt.Println("initializing new directory...")

	// TODO: How do I set the source when installed on a new computer?
	src := "/Users/jmckee3/go/src/github.com/jmckee46/deployer/new-directory"

	// Get current working directory
	dst, err := os.Getwd()
	if err != nil {
		logger.Panic("initNewDirectory", flaw.From(err))
	}

	// modified version of copy directory
	files, err := ioutil.ReadDir(src)
	if err != nil {
		logger.Panic("initNewDirectory", flaw.From(err))
	}

	if len(files) > 0 {
		for _, file := range files {
			newSrc := filepath.Join(src, file.Name())
			newDst := filepath.Join(dst, file.Name())
			if file.IsDir() {
				osfuncs.CopyDirectory(newSrc, dst)
				continue
			}
			_, copyErr := osfuncs.CopyFile(newSrc, newDst)
			if copyErr != nil {
				logger.Panic("initNewDirectory", copyErr.String())
			}
		}
	}

	flawErr := osfuncs.CopyDirectory(src, dst)
	if flawErr != nil {
		logger.Panic("initNewDirectory", flawErr)
	}

	flawErr = updateImports(dst)
	if flawErr != nil {
		logger.Panic("initNewDirectory", flawErr)
	}
}

func updateImports(curDir string) flaw.Flaw {
	// find all the go files
	var files []string

	err := filepath.Walk(curDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return flaw.From(err)
	}

	// replace /myAppTest/ with /curDir/
	base := filepath.Base(curDir)
	for _, file := range files {
		read, err := ioutil.ReadFile(file)
		if err != nil {
			return flaw.From(err)
		}

		newDir := "/" + base + "/"

		newContents := strings.Replace(string(read), "/myAppTest/", newDir, -1)

		err = ioutil.WriteFile(file, []byte(newContents), 0)
		if err != nil {
			return flaw.From(err)
		}
	}

	return nil
}

// Create directories
// 	dirsToMake := []string{
// 		"images-to-deploy",
// 		"docker-compose-only-images/load-balancer",
// 		"docker-compose-only-images/mocks",
// 	}

// 	for _, dir := range dirsToMake {
// 		flaw := createDir(dst, dir)
// 		if flaw != nil {
// 			logger.Panic("initNewDirectory", flaw)
// 		}
// 	}

// 	// Copy directories
// 	dirsToCopy := []string{
// 		"docker-compose-only-images",
// 	}

// 	for _, dir := range dirsToCopy {
// 		newSrc := src + dir
// 		newDst := dst + dir
// 		flaw := osfuncs.CopyDirectory(newSrc, newDst)
// 		if flaw != nil {
// 			logger.Panic("initNewDirectory", flaw)
// 		}
// 	}

// 	// Copy base level files
// 	filesToCopy := []string{
// 		".env",
// 		"docker-compose.yaml",
// 	}

// 	for _, file := range filesToCopy {
// 		newSrc := src + file
// 		newDst := dst + file
// 		_, flaw := osfuncs.CopyFile(newSrc, newDst)
// 		if flaw != nil {
// 			logger.Panic("initNewDirectory", flaw)
// 		}
// 	}
// }

// func createDir(dstBase, name string) *flaw.Error {
// 	newDst := dstBase + name
// 	err := os.MkdirAll(newDst, 0755)
// 	if err != nil {
// 		return flaw.From(err)
// 	}

// 	return nil
// }
