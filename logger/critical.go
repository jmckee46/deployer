package logger

import (
	"github.com/jmckee46/deployer/color"
)

// Critical sends a critical level log message
func Critical(tag string, message interface{}) {
	send(
		"critical",
		color.LightRed,
		tag,
		message,
	)
}
