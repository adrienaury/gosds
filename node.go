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
	Index() int
	Key() string
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
	IsArray() bool
	IsObject() bool
	IsRoot() bool
	AsKeyed() Keyed
	AsIndexed() Indexed
	AsArray() Array
	AsObject() Object
	AsRoot() Root
}
