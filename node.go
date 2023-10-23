package gosds

type Node interface {
	Parented
	Container
	Castable
	Marshaler
}

type Parented interface {
	Root() Node
	Parent() Node
}

type Container interface {
	Get() any
	Set(val any)
	Remove()
	Primitive() any
	Exist() bool
}

type Castable interface {
	IsKeyed() bool
	IsIndexed() bool
	AsKeyed() Keyed
	AsIndexed() Indexed
}
