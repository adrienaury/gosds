package gosds

type Object interface {
	Keyed
	Indexed
}

type object struct {
	capacity int
}

func newObjectWithCapacity(capacity int) *value {
	return newValue(
		&object{capacity},
	)
}
