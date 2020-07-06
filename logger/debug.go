package logger

import (
	"github.com/jmckee46/deployer/color"
)

// Debug sends a debug level log message
func Debug(tag string, message interface{}) {
	if DeLogDebug != "true" {
		return
	}

	send("debug", color.LightBlack, tag, message)
}
