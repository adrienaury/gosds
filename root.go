package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Root interface {
	Node
}

type root struct {
	value Node
}

func newRoot(val Node) *root {
	return &root{
		value: val,
	}
}

func (r *root) Get() any {
	return r.value
}

func (r *root) Set(val any) {
	switch typedValue := val.(type) {
	case *object:
		typedValue.parent = r
		typedValue.index = 0
		r.value = typedValue
	case *array:
		typedValue.parent = r
		typedValue.index = 0
		r.value = typedValue
	case *value:
		typedValue.parent = r
		typedValue.index = 0
		r.value = typedValue
	case object:
		typedValue.parent = r
		typedValue.index = 0
		r.value = &typedValue
	case array:
		typedValue.parent = r
		typedValue.index = 0
		r.value = &typedValue
	case value:
		typedValue.parent = r
		typedValue.index = 0
		r.value = &typedValue
	default:
		panic("not accepted")
	}
}

func (r *root) Primitive() any {
	return r.value
}

func (r *root) AsObject() (Object, bool)   { return nil, false } //nolint:ireturn
func (r *root) AsArray() (Array, bool)     { return nil, false } //nolint:ireturn
func (r *root) AsValue() (Value, bool)     { return nil, false } //nolint:ireturn
func (r *root) AsIndexed() (Indexed, bool) { return nil, false } //nolint:ireturn
func (r *root) MustObject() Object         { return nil }        //nolint:ireturn
func (r *root) MustArray() Array           { return nil }        //nolint:ireturn
func (r *root) MustValue() Value           { return nil }        //nolint:ireturn
func (r *root) MustIndexed() Indexed       { return nil }        //nolint:ireturn

func (r *root) Parent() Node { return nil } //nolint:ireturn
func (r *root) Index() int   { return 0 }

func (r *root) MarshalEncode(w *jwriter.Writer) { r.value.MarshalEncode(w) }
func (r *root) MarshalWrite(w io.Writer) error  { return r.value.MarshalWrite(w) } //nolint:wrapcheck