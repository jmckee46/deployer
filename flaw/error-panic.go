package flaw

// Panic panic with killer tracing
func (flawErr *Error) Panic() {
	println("\n\n" + flawErr.String() + "\n\n")
	panic("panic requested!")
}
