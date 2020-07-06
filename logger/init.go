package logger

import (
	"bytes"
	"os"

	"github.com/jmckee46/deployer/color"
	"github.com/jmckee46/deployer/serializers"
)

var colorizer func(string, []byte) *bytes.Buffer
var serializer func(interface{}) []byte

// DeLogDebug is a variable set by environment variable
var DeLogDebug string

// DeLogColorization is a variable set by environment variable
var DeLogColorization string

// DeLogSerialization is a variable set by environment variable
var DeLogSerialization string

// DeGitBranch is a variable set by environment variable
var DeGitBranch string

// DeGitSha is a variable set by environment variable
var DeGitSha string

func init() {
	DeLogDebug = os.Getenv("DE_LOG_DEBUG_MESSAGES")
	DeLogColorization = os.Getenv("DE_LOG_COLORIZATION")
	DeLogSerialization = os.Getenv("DE_LOG_SERIALIZATION")
	DeGitBranch = os.Getenv("DE_GIT_BRANCH")
	DeGitSha = os.Getenv("DE_GIT_SHA")

	if DeLogColorization == "true" {
		colorizer = color.Colorize
	} else {
		colorizer = color.Discolorize
	}

	switch DeLogSerialization {
	case "json-pretty":
		serializer = serializers.JSONPrettySerializer
	case "json-compact":
		serializer = serializers.JSONCompactSerializer
	default:
		serializer = serializers.JSONCompactSerializer
	}
}
