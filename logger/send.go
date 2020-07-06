package logger

func send(severity string, ansiPrefix string, tag string, message interface{}) {
	anEntry := &entry{
		ansiPrefix: ansiPrefix,
		GitBranch:  DeGitBranch,
		GitSha:     DeGitSha,
		Message:    message,
		Severity:   severity,
		Tag:        tag,
	}

	serializeAndWrite(anEntry)
}
