package stamps

import (
	"time"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/serializers"
)

type Timestamp struct {
	*time.Time
}

// New returns the current time
func New() *Timestamp {
	stamp := time.Now().UTC()

	return &Timestamp{&stamp}
}

// NewAt returns the current time
func NewAt(time *time.Time) *Timestamp {

	return &Timestamp{time}
}

func (stamp *Timestamp) Scan(src interface{}) error {
	source, ok := src.([]byte)

	if !ok {
		return flaw.New("wrong type").Wrap("cannot Scan")
	}

	flawError := serializers.JSONUnmarshalTag("json", source, stamp)

	if flawError != nil {
		return flawError.Wrap("cannot Scan")
	}

	return nil
}
