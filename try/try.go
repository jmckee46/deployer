package try

import (
	"time"

	"github.com/jmckee46/deployer/flaw"
)

type Try struct {
	Start   time.Time     `json:"start"`
	Max     uint          `json:"max"`
	Tries   uint          `json:"tries"`
	Elapsed time.Duration `json:"elapsed"`
	Error   flaw.Flaw     `json:"error"`
}
