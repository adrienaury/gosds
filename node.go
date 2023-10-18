package gosds

type Node interface {
	Parent() Node
	Value() any

	AsObject() (Object, bool)
	AsArray() (Array, bool)

	// Primitive returns a representation of the node with following types :
	// - objects as map[string]any
	// - arrays as []any
	// - values as string, float64, bool or nil interface
	Primitive() any
}
