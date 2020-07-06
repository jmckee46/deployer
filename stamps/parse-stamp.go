package stamps

import (
	"time"

	"github.com/jmckee46/deployer/flaw"
)

func ParseStamp(stamp string, format string) (*Timestamp, flaw.Flaw) {
	parsed, err := time.ParseInLocation(format, stamp, time.UTC)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot ParseStamp")
	}

	return &Timestamp{&parsed}, nil
}
