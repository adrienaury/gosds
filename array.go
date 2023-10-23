package gosds

type Array interface {
	Indexed
	AppendValue(value any)
}

type array struct {
	capacity int
}

func newArrayWithCapacity(capacity int) *value {
	return newValue(
		&array{capacity},
	)
}
