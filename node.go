package gosds

type Node interface {
	Parent() Node
	Value() any
}
