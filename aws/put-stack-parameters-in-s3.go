package awsfuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jmckee46/deployer/encrypt"
	"github.com/jmckee46/deployer/flaw"
)

// PutStackParametersInS3 puts the stackParametersState in S3
func PutStackParametersInS3(state *State) flaw.Flaw {
	fmt.Println("  putting stack parameters in S3...")

	// marshal the stack parameters
	b, err := json.Marshal(state.StackParametersStackState)
	if err != nil {
		return flaw.From(err)
	}

	// encrypt the stack parameters
	keytext := os.Getenv("DE_PASS_PHRASE")
	key := encrypt.StringToByte(keytext)

	cyphertext, err := encrypt.Encrypt(b, key)
	if err != nil {
		return flaw.From(err)
	}

	// save stack parameters to file
	stackParametersPath := filepath.Join(
		os.Getenv("DE_ARTIFACTS_PATH"),
		os.Getenv("DE_GIT_BRANCH"),
		"stack-parameters",
	)

	err = ioutil.WriteFile(stackParametersPath, cyphertext, 0755)
	if err != nil {
		return flaw.From(err)
	}

	// upload stack parameters to S3
	s3path := filepath.Join(
		os.Getenv("DE_GIT_BRANCH"),
		"stack-parameters",
		os.Getenv("DE_GIT_SHA"),
	) + ".stackparameters"

	err = PutFileInS3(state, s3path, stackParametersPath)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
