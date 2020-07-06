package logger

import "github.com/jmckee46/deployer/halt"

// Panic sends a panic level log message
func Panic(tag string, message interface{}) {
	Critical(tag, message)

	halt.PanicWith("halt requested via logger.Panic")
}
