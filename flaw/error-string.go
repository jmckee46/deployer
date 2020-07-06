package flaw

import "strings"

// String implements fmt.Stringer
func (err *Error) String() string {
	return strings.TrimSpace(
		"message trace\n" +
			"-----------\n" +
			err.messageTrace() +
			"\n" +
			"stack trace\n" +
			"-----------\n" +
			err.stackTrace(),
	)
}
