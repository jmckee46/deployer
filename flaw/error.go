package flaw

// Error contains the failure chain stacks
type Error struct {
	MessageTrace []messageTrace `json:"message-trace"`
	StackTrace   []frame        `json:"stack-trace"`
}
