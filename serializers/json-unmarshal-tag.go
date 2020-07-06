package serializers

import (
	jsoniter "github.com/halorium/json-iterator"
	"github.com/jmckee46/deployer/flaw"
)

// JSONUnmarshalTag is a type of deserializer
func JSONUnmarshalTag(tag string, data []byte, entity interface{}) flaw.Flaw {
	err := jsoniter.Config{Tag: tag}.Froze().Unmarshal(data, entity)

	if err != nil {
		return flaw.From(err).Wrap("cannot JSONUnmarshalTag")
	}

	return nil
}
