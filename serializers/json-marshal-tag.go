package serializers

import (
	jsoniter "github.com/halorium/json-iterator"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// JSONMarshalTag is a type of serializer
func JSONMarshalTag(tag string, entity interface{}) []byte {
	serialized, err := jsoniter.Config{Tag: tag}.Froze().Marshal(entity)

	if err != nil {
		halt.Panic(flaw.From(err))
	}

	return append(serialized, "\n"...)
}
