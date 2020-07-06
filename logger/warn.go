package logger

import "github.com/jmckee46/deployer/color"

// Warn sends a warning level log message
func Warn(tag string, message interface{}) {
	send("warn", color.LightMagenta, tag, message)
}
