package halt

import "github.com/jmckee46/deployer/flaw"

func PanicWith(message string) {
	panic(flaw.New(message))
}
