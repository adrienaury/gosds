package gosds

type Builder interface {
	AddKey(key string) error
	AddValue(val any) error
	StartObject() error
	StartArray() error
	EndObjectOrArray() error
	Build() Node
}
