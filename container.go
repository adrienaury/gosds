package gosds

type Container interface {
	Get() any
	Set(val any)
	Primitive() any
	Exist() bool
}
