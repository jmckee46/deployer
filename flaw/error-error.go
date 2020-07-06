package flaw

// Error implements error
func (flawErr *Error) Error() string {
	firstErrorIndex := len(flawErr.MessageTrace) - 1

	return flawErr.MessageTrace[firstErrorIndex].Message
}
