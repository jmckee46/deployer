package loadBalancer

import (
	"fmt"
	"testing"
)

// Test doesn't run from load-balancer directory as is, need to change
// file paths in create-cert.go - remove both "mocks/load-balancer/"
func TestCreateCert(t *testing.T) {
	fmt.Println("testing")

	CreateCert()

	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
