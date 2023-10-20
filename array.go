package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Array interface {
	Node

	Indexed

	AppendValue(value any)

	// PrimitiveArray returns a representation of the array as []any
	PrimitiveArray() []any
}

type array struct {
	values []Node

	parent Node
	index  int
}

func NewArray() Array { //nolint:ireturn
	return newArray()
}

func NewArrayWithCapacity(capacity int) Array { //nolint:ireturn
	return newArrayWithCapacity(capacity)
}

func newArray() *array {
	return &array{
		values: []Node{},
		parent: nil,
		index:  0,
	}
}

func newArrayWithCapacity(capacity int) *array {
	return &array{
		values: make([]Node, 0, capacity),
		parent: nil,
		index:  0,
	}
}

func (a *array) Parent() Node { //nolint:ireturn
	return a.parent
}

func (a *array) Index() int {
	return a.index
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

func (a *array) AsIndexed() (Indexed, bool) { //nolint:ireturn
	return a, true
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

func (a *array) MustIndexed() Indexed { //nolint:ireturn
	return a
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

func (a *array) NodeAtIndex(index int) Node { //nolint:ireturn
	return a.values[index]
}

func (a *array) SetValueAtIndex(index int, val any) {
	a.values = set(a.values, val, index, a)
}

func (a *array) RemoveValueAtIndex(index int) {
	a.values = append(a.values[:index], a.values[index+1:]...)
}

func (a *array) AppendValue(val any) {
	a.values = add(a.values, val, a)
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
