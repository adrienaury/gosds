package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Root interface {
	Parented
	Container
	Castable
	JSONMarshaller
}

type root struct {
	value Node
}

func newRoot(val Node) *root {
	root := &root{} //nolint:exhaustruct
	root.Set(val)

	return root
}

func (r *root) Get() any {
	return r.value.Get()
}

func (r *root) Set(val any) {
	switch typedValue := val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		value := newValue(val)
		value.root = r
		value.parent = nil
		value.index = 0
		r.value = value
	case *object:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = typedValue
	case *array:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = typedValue
	case *value:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = typedValue
	case object:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = &typedValue
	case array:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = &typedValue
	case value:
		typedValue.root = r
		typedValue.parent = nil
		typedValue.index = 0
		r.value = &typedValue
	default:
		panic("not accepted")
	}
}

func (r *root) Primitive() any {
	return r.value.Primitive()
}

func (r *root) Exist() bool                { return r.value.Exist() }
func (r *root) AsObject() (Object, bool)   { return r.value.AsObject() }    //nolint:ireturn
func (r *root) AsArray() (Array, bool)     { return r.value.AsArray() }     //nolint:ireturn
func (r *root) AsValue() (Value, bool)     { return r.value.AsValue() }     //nolint:ireturn
func (r *root) AsIndexed() (Indexed, bool) { return r.value.AsIndexed() }   //nolint:ireturn
func (r *root) AsRoot() (Root, bool)       { return r, true }               //nolint:ireturn
func (r *root) MustObject() Object         { return r.value.MustObject() }  //nolint:ireturn
func (r *root) MustArray() Array           { return r.value.MustArray() }   //nolint:ireturn
func (r *root) MustValue() Value           { return r.value.MustValue() }   //nolint:ireturn
func (r *root) MustIndexed() Indexed       { return r.value.MustIndexed() } //nolint:ireturn
func (r *root) MustRoot() Root             { return r }                     //nolint:ireturn

func (r *root) Parent() Node { return nil } //nolint:ireturn
func (r *root) Index() int   { return 0 }

func (r *root) MarshalEncode(w *jwriter.Writer) { r.value.MarshalEncode(w) }
func (r *root) MarshalWrite(w io.Writer) error  { return r.value.MarshalWrite(w) } //nolint:wrapcheck
