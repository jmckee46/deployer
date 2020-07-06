package interfaces

type Entity interface {
	Prefix() string
	ID() string
	ETag() string
	CreatedAt() string
	EffectiveAt() string
	Header() []byte
	Body() []byte
}
