package interfaces

import (
	"fmt"
)

// Flaw is our interface to flaws (wrapped errors)
type Flaw interface {
	error
	fmt.Stringer
	Panic()
	Wrap(message string) Flaw
}
