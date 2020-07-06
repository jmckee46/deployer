package git

import (
	"fmt"
	"testing"
)

func TestCurrentSha(t *testing.T) {
	sha := CurrentSha()

	fmt.Println("the current sha is:", sha)

	// if daysNotified != 1 {
	// 	t.Errorf("got %d instead of 1", daysNotified)
	// }
}
