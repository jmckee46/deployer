package tlsDeployer

import (
	"fmt"

	"github.com/jmckee46/deployer/certbot"
	"github.com/jmckee46/deployer/flaw"
)

func ManageFiles() flaw.Flaw {
	fmt.Println("managing tls files...")

	state := newState()

	if tlsFilesInS3Fresh(state) {
		err := getTLSFilesFromS3(state)
		if err != nil {
			return err
		}

		return nil
	}

	err := certbot.GetTLSFilesFromLetsencrypt()
	if err != nil {
		return err
	}

	err = certbot.CopyTLSFilesToTLSDirectory()
	if err != nil {
		return err
	}

	err = PutTLSFilesInACMAndARNInFile()
	if err != nil {
		return err
	}

	err = PutTLSFilesInS3()
	if err != nil {
		return err
	}

	return nil
}
