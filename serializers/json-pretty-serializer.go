package serializers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
)

// JSONPrettySerializer is a type of serializer
func JSONPrettySerializer(object interface{}) []byte {
	buffer := bytes.NewBuffer(nil)

	encoder := json.NewEncoder(buffer)

	encoder.SetIndent("", "  ")

	err := encoder.Encode(object)

	if err != nil {
		fmt.Println(err)
		halt.Panic(flaw.From(err))
	}

	return buffer.Bytes()
}
