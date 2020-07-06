package serializers

import (
	"encoding/json"
	"io"

	"github.com/jmckee46/deployer/flaw"
)

// JSONDeserializer is a type of deserializer
func JSONDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) flaw.Flaw {
	defer readCloser.Close()

	err := json.NewDecoder(readCloser).Decode(deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot JSONDeserializer")
	}

	return nil
}
