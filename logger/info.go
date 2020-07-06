package logger

import "github.com/jmckee46/deployer/color"

// Info sends an info level log message
func Info(tag string, message interface{}) {
	send("info", color.LightYellow, tag, message)
}
