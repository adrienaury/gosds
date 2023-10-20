package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Value interface {
	Node

	Value() any

	Set(val any)

	// Int64() (int64, bool)
	// MustInt64() int64
}

type value struct {
	parent Node
	index  int
	value  any
}

func newValue(val any) *value {
	return &value{
		parent: nil,
		index:  0,
		value:  val,
	}
}

func (v *value) Parent() Node { //nolint:ireturn
	return v.parent
}

func (v *value) Index() int {
	return v.index
}

func (v *value) Value() any {
	return v.value
}

func (v *value) AsObject() (Object, bool) { //nolint:ireturn
	return nil, false
}

func (v *value) AsArray() (Array, bool) { //nolint:ireturn
	return nil, false
}

func (v *value) AsValue() (Value, bool) { //nolint:ireturn
	return v, true
}

func (v *value) MustObject() Object { //nolint:ireturn
	return nil
}

func (v *value) MustArray() Array { //nolint:ireturn
	return nil
}

func (v *value) MustValue() Value { //nolint:ireturn
	return v
}

func (v *value) Primitive() any {
	return v.value
}

func (v *value) Set(val any) {
	switch val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		v.value = val
	default:
		panic("not accepted")
	}
}

func (v *value) MarshalEncode(output *jwriter.Writer) { //nolint:cyclop
	switch typedValue := v.value.(type) {
	case string:
		output.String(typedValue)
	case int:
		output.Int(typedValue)
	case int64:
		output.Int64(typedValue)
	case int32:
		output.Int32(typedValue)
	case int16:
		output.Int16(typedValue)
	case int8:
		output.Int8(typedValue)
	case uint:
		output.Uint(typedValue)
	case uint64:
		output.Uint64(typedValue)
	case uint32:
		output.Uint32(typedValue)
	case uint16:
		output.Uint16(typedValue)
	case uint8:
		output.Uint8(typedValue)
	case float64:
		output.Float64(typedValue)
	case float32:
		output.Float32(typedValue)
	case bool:
		output.Bool(typedValue)
	case Number:
		output.RawString(typedValue.String())
	case nil:
		output.RawString("null")
	}
}

func (v *value) MarshalWrite(output io.Writer) error {
	return MarshalWrite(v, output)
}
