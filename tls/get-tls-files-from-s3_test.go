package tlsDeployer

import (
	"fmt"
	"testing"
)

// TestGetTLSFilesFromS3 cannot run without removing the destination "tls" as the
// test runs in tls directory and not deployer directory
func TestGetTLSFilesFromS3(t *testing.T) {
	t.Skip("skipping - cannot run from tls directory")
	state := newState()

	err := getTLSFilesFromS3(state)

	if err != nil {
		fmt.Println("err:", err)
	}
}
