package image

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/git"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/osfuncs"
)

// Build performs a docker build on the path passed
func Build(path string) flaw.Flaw {
	fmt.Printf("  building                       %s...\n", path)

	currentSha := git.CurrentSha()
	baseName, flawErr := getBaseName()
	if flawErr != nil {
		logger.Panic("image-build err, are you running deployer from the base directory?", flawErr.String())
	}
	imageID := baseName + "/" + path
	imageTag := imageID + ":" + currentSha
	imageTag = strings.TrimSpace(imageTag)
	imageTag2 := imageID + ":latest"
	imageTag2 = strings.TrimSpace(imageTag2)

	// docker build the image with tag
	cmd := exec.Command("docker", "build", "-t", imageTag, ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Dir = path
	err := cmd.Run()
	// fmt.Printf("output: %s\n", out.String())
	if err != nil {
		// return flaw.From(err).Wrap("image build failed").Wrap(out.String())
		return flaw.New(out.String())
	}

	// also tag as latest
	_, err = exec.Command("docker", "tag", imageTag, imageTag2).Output()
	if err != nil {
		return flaw.From(err).Wrap("image build failed")
	}

	return nil
}

// getBaseName gets the base name to include with tagging. For this to work properly, deployer
// should be run from the root directory and getBaseName checks for images-to-deploy to ensure this.
func getBaseName() (string, flaw.Flaw) {
	cmd := exec.Command("pwd")
	workingDir, err := cmd.CombinedOutput()
	if err != nil {
		return "", flaw.From(err).Wrap("get base name failed")
	}

	fullBasePath := fmt.Sprintf("%s", workingDir)
	fullBasePath = strings.TrimSpace(fullBasePath)

	itdPath := fullBasePath + "/images-to-deploy"

	exists, err := osfuncs.Exists(itdPath)
	if !exists {
		return "", flaw.From(err).Wrap("get base name failed")
	}

	baseName := filepath.Base(fullBasePath)
	lowercasebase := strings.ToLower(baseName)

	return lowercasebase, nil
}
