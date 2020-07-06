package flaw

// New creates a new flaw error from a string
func New(message string) *Error {
	return create(message)
}
