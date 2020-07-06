package serializers

import "io"

// Deserializer defines a generic deserializer
type Deserializer func(io.ReadCloser, interface{}) error

// Serializer defines a generic serializer
type Serializer func(interface{}) []byte
