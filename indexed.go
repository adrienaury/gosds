package gosds

type Indexed interface {
	NodeAtIndex(index int) Node
	ValueAtIndex(index int) any
	SetValueAtIndex(index int, value any)
	RemoveValueAtIndex(index int)
	Size() int
	PrimitiveArray() []any
}
