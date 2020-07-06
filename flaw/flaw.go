package flaw

import "fmt"

// Flaw is a kind of error with great tracing
type Flaw interface {
	error
	fmt.Stringer
	Panic()
	Wrap(message string) Flaw
}
