package serializers

import (
	"encoding/xml"
	"io"

	"github.com/jmckee46/deployer/flaw"
)

// XMLDeserializer is a type of deserializer
func XMLDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) flaw.Flaw {
	defer readCloser.Close()

	err := xml.NewDecoder(readCloser).Decode(deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot XMLDeserializer")
	}

	return nil
}
