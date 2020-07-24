package tlsDeployer

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/osfuncs"
)

// TestGetTLSFilesFromS3 cannot run without removing the destination "tls" as the
// test runs in tls directory and not deployer directory
func TestGetTLSFilesFromS3(t *testing.T) {
	// t.Skip("skipping - cannot run from tls directory")

	setUp()
	defer tearDown()

	state := newState()

	err := getTLSFilesFromS3(state)

	if err != nil {
		fmt.Println("err:", err)
	}
}

func setUp() {
	// since tests run in the tls folder vice deployer...
	err := os.MkdirAll("tls/files", 0755)
	if err != nil {
		logger.Panic("TestGetTLSFilesFromS3", err)
	}
}

func tearDown() {
	osfuncs.DeleteDirAndFiles("tls")
}
