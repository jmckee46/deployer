package flaw

// From creates a flaw error from an error
func From(err error) *Error {
	if err == nil {
		return nil
	}

	return create(err.Error())
}
