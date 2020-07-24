package main

import (
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

func Cleanup() flaw.Flaw {
	err := osfuncs.DeleteDirAndFiles(".aws")
	if err != nil {
		return err
	}

	err = osfuncs.DeleteDirAndFiles("certbot/files")
	if err != nil {
		return err
	}

	err = osfuncs.DeleteDirAndFiles("certbot/log-files")
	if err != nil {
		return err
	}

	err = osfuncs.DeleteDirAndFiles("tls/files")
	if err != nil {
		return err
	}

	return nil
}
