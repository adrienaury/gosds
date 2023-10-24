package gosds

type Node interface {
	Parented
	Container
	Castable
	Marshaler
}

type Parented interface {
	Root() Root
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
	IsObject() bool
	IsArray() bool
	IsRoot() bool
	AsKeyed() Keyed
	AsIndexed() Indexed
	AsObject() Object
	AsArray() Array
	AsRoot() Root
}
