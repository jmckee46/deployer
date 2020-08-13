package imagesToDeploy

import (
	"fmt"
	"testing"
)

func TestHttpSpecBuild(t *testing.T) {
	err := BuildMigrator()

	if err != nil {
		fmt.Println("err from test:", err.String())
	}

	// err := gofuncs.Build("images-to-deploy/health-check")

	// if err != nil {
	// 	fmt.Println(err.String())
	// }

	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
