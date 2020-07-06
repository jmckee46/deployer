package logger

type emptyMessage struct{}

// EmptyMessage logs and empty message
func EmptyMessage() emptyMessage {
	return emptyMessage{}
}
