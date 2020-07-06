package try

import (
	"time"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

// retries 3 times after waiting 1 second, then 9 seconds, then 14 seconds

func (t *Try) Again(tag string, flawError flaw.Flaw) bool {
	t.Elapsed = time.Since(t.Start)
	t.Error = flawError
	t.Tries++

	if t.Tries == t.Max {
		logger.Critical(tag, t)

		return false
	}

	logger.Warn(tag, t)

	seconds := time.Duration(t.Tries)

	time.Sleep(seconds * seconds * time.Second)

	return true
}
