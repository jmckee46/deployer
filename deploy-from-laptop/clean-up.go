package main

import (
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/osfuncs"
)

func Cleanup() flaw.Flaw {
	// ok to fail if not there
	osfuncs.DeleteDirAndFiles(".aws")

	// ok to fail if not there
	osfuncs.DeleteDirAndFiles("certbot/files")

	// ok to fail if not there
	osfuncs.DeleteDirAndFiles("certbot/log-files")

	// ok to fail if not there
	osfuncs.DeleteDirAndFiles("tls/files")

	return nil
}
