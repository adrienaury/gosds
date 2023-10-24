package gosds

import (
	"io"
)

type Root interface {
	Parented
	Container
	Castable
	Marshaler
}

type root struct {
	value Node
}

func newRoot(val Node) *root {
	root := &root{} //nolint:exhaustruct
	root.Set(val)

	return root
}

func (r *root) Set(val any) {
	switch typedValue := accept(val, "", 0, nil).(type) {
	case *object:
		typedValue.root = r
		r.value = typedValue
	case *array:
		typedValue.root = r
		r.value = typedValue
	case *value:
		typedValue.root = r
		r.value = typedValue
	case *placeholder:
		typedValue.root = r
		r.value = typedValue
	default:
		panic("not accepted")
	}
}

func (r *root) Root() Root         { return r }   //nolint:ireturn
func (r *root) Parent() Node       { return nil } //nolint:ireturn
func (r *root) Index() int         { return 0 }
func (r *root) Key() string        { return "" }
func (r *root) Get() any           { return r.value.Get() }
func (r *root) Remove()            { r.value = nil }
func (r *root) Primitive() any     { return r.value.Primitive() }
func (r *root) Exist() bool        { return r.value.Exist() }
func (r *root) IsKeyed() bool      { return r.value.IsKeyed() }
func (r *root) IsIndexed() bool    { return r.value.IsIndexed() }
func (r *root) IsObject() bool     { return r.value.IsObject() }
func (r *root) IsArray() bool      { return r.value.IsArray() }
func (r *root) IsRoot() bool       { return true }
func (r *root) AsKeyed() Keyed     { return r.value.AsKeyed() }   //nolint:ireturn
func (r *root) AsIndexed() Indexed { return r.value.AsIndexed() } //nolint:ireturn
func (r *root) AsObject() Object   { return r.value.AsObject() }  //nolint:ireturn
func (r *root) AsArray() Array     { return r.value.AsArray() }   //nolint:ireturn
func (r *root) AsRoot() Root       { return r }                   //nolint:ireturn

func (r *root) MarshalEncode(w Encoder)        { r.value.MarshalEncode(w) }
func (r *root) MarshalWrite(w io.Writer) error { return r.value.MarshalWrite(w) } //nolint:wrapcheck
