package flaw

import (
	"runtime"
)

// Wrap wraps one flaw with another creating a failure chain
func (flawErr *Error) Wrap(message string) Flaw {
	_, pathname, line, ok := runtime.Caller(1)

	if !ok {
		panic("not ok")
	}

	messageTraceStruct := messageTrace{
		Message:  message,
		Pathname: stripPathname(pathname),
		Line:     line,
	}

	flawErr.MessageTrace = append(
		[]messageTrace{messageTraceStruct}, flawErr.MessageTrace...,
	)

	return flawErr
}
