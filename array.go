package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Array interface {
	Node

	ValueAtIndex(index int) any
	SetValueAtIndex(index int, value any)
	AppendValue(value any)

	Size() int

	// PrimitiveArray returns a representation of the array as []any
	PrimitiveArray() []any
}

type array struct {
	values []Node

	parent Node
}

func newArray(parent Node) Array { //nolint:ireturn
	return &array{
		values: []Node{},
		parent: parent,
	}
}

func newArrayWithCapacity(parent Node, capacity int) Array { //nolint:ireturn
	return &array{
		values: make([]Node, 0, capacity),
		parent: parent,
	}
}

func (a *array) Parent() Node { //nolint:ireturn
	return a.parent
}

func (a *array) Value() any {
	return a
}

func (a *array) AsObject() (Object, bool) { //nolint:ireturn
	return nil, false
}

func (a *array) AsArray() (Array, bool) { //nolint:ireturn
	return a, true
}

func (a *array) AsValue() (Value, bool) { //nolint:ireturn
	return nil, false
}

func (a *array) MustObject() Object { //nolint:ireturn
	return nil
}

func (a *array) MustArray() Array { //nolint:ireturn
	return a
}

func (a *array) MustValue() Value { //nolint:ireturn
	return nil
}

func (a *array) Primitive() any {
	result := make([]any, len(a.values))

	for index, val := range a.values {
		result[index] = val.Primitive()
	}

	return result
}

func (a *array) PrimitiveArray() []any {
	return a.Primitive().([]any) //nolint:forcetypeassert
}

func (a *array) ValueAtIndex(index int) any {
	return a.values[index].Value()
}

func (a *array) SetValueAtIndex(index int, value any) {
	switch typedValue := value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		a.values[index] = newValue(a, value)
	case Node:
		a.values[index] = typedValue
	default:
		panic("not accepted")
	}
}

func (a *array) AppendValue(value any) {
	switch typedValue := value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		a.values = append(a.values, newValue(a, value))
	case Node:
		a.values = append(a.values, typedValue)
	default:
		panic("not accepted")
	}
}

func (a *array) Size() int {
	return len(a.values)
}

func (a *array) MarshalEncode(output *jwriter.Writer) {
	output.RawByte('[')

	for index, node := range a.values {
		if index > 0 {
			output.RawByte(',')
		}

		node.MarshalEncode(output)
	}

	output.RawByte(']')
}

func (a *array) MarshalWrite(output io.Writer) error {
	return MarshalWrite(a, output)
}
