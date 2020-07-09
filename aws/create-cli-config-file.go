package awsfuncs

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// CreateCliConfigFile creates the config file for aws
func CreateCliConfigFile() flaw.Flaw {
	fmt.Println("creating aws config file...")
	err := os.MkdirAll(".aws", 0755)
	if err != nil {
		return flaw.From(err)
	}

	f, err := os.Create(".aws/config")
	if err != nil {
		return flaw.From(err)
	}

	_, err = f.WriteString("[default]\n")
	if err != nil {
		f.Close()
		return flaw.From(err)
	}

	_, err = f.WriteString("aws_access_key_id=$AWS_ACCESS_KEY_ID\n")
	if err != nil {
		f.Close()
		return flaw.From(err)
	}

	_, err = f.WriteString("aws_secret_access_key=$AWS_SECRET_ACCESS_KEY\n")
	if err != nil {
		f.Close()
		return flaw.From(err)
	}

	err = f.Close()
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
