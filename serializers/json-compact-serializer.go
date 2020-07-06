package serializers

import (
	"encoding/json"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// JSONCompactSerializer is a type of serializer
func JSONCompactSerializer(object interface{}) []byte {
	serialized, err := json.Marshal(object)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return append(serialized, "\n"...)
}
