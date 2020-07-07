package serializers

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/jmckee46/deployer/flaw"
)

// CSVDeserializer is a type of deserializer
func CSVDeserializer(readCloser io.ReadCloser, deserializedEntity interface{}) flaw.Flaw {
	defer readCloser.Close()

	err := gocsv.Unmarshal(readCloser, deserializedEntity)

	if err != nil {
		return flaw.From(err).Wrap("cannot CSVDeserializer")
	}

	return nil
}
