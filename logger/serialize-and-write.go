package logger

import (
	"os"

	"github.com/jmckee46/deployer/flaw"
)

func serializeAndWrite(entry *entry) {
	serialized := serializer(entry)

	serializedNoNewline := serialized[:len(serialized)-1]

	buffer := colorizer(entry.ansiPrefix, serializedNoNewline)

	buffer.WriteString("\n")

	_, err := buffer.WriteTo(os.Stdout)

	if err != nil {
		panic(flaw.From(err).Wrap("cannot serializeAndWrite"))
	}
}
