package main

import (
	"fmt"
	"testing"

	"github.com/jmckee46/deployer/tls"
)

func TestCurrentSha(t *testing.T) {
	// err := awsfuncs.GetTlsFilesFromS3()
	err := tlsDeployer.ManageFiles()
	if err != nil {
		fmt.Println(err.String())
	}
	// fmt.Println("calling base unit test")
	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
