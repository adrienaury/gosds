package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Object interface {
	Node

	NodeForKey(key string) Node

	ValueForKey(key string) (any, bool)
	SetValueForKey(key string, value any)
	RemoveValueForKey(key string)

	Keys() []string

	// PrimitiveObject returns a representation of the object as map[string]any
	PrimitiveObject() map[string]any
}

type object struct {
	inner map[string]Node
	keys  []string

	parent Node
}

func newObject(parent Node) Object { //nolint:ireturn
	return &object{
		inner:  map[string]Node{},
		keys:   []string{},
		parent: parent,
	}
}

func newObjectWithCapacity(parent Node, capacity int) Object { //nolint:ireturn
	return &object{
		inner:  make(map[string]Node, capacity),
		keys:   make([]string, 0, capacity),
		parent: parent,
	}
}

func (o *object) Parent() Node { //nolint:ireturn
	return o.parent
}

func (o *object) Value() any {
	return o
}

func (o *object) AsObject() (Object, bool) { //nolint:ireturn
	return o, true
}

func (o *object) AsArray() (Array, bool) { //nolint:ireturn
	return nil, false
}

func (o *object) AsValue() (Value, bool) { //nolint:ireturn
	return nil, false
}

func (o *object) MustObject() Object { //nolint:ireturn
	return o
}

func (o *object) MustArray() Array { //nolint:ireturn
	return nil
}

func (o *object) MustValue() Value { //nolint:ireturn
	return nil
}

func (o *object) Primitive() any {
	result := make(map[string]any, len(o.keys))

	for key, val := range o.inner {
		result[key] = val.Primitive()
	}

	return result
}

func (o *object) PrimitiveObject() map[string]any {
	return o.Primitive().(map[string]any) //nolint:forcetypeassert
}

func (o *object) NodeForKey(key string) Node { //nolint:ireturn
	return o.inner[key]
}

func (o *object) ValueForKey(key string) (any, bool) {
	node, has := o.inner[key]

	return node.Value(), has
}

func (o *object) SetValueForKey(key string, value any) {
	o.keys = append(o.keys, key)
	switch typedValue := value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		o.inner[key] = newValue(o, value)
	case Node:
		o.inner[key] = typedValue
	default:
		panic("not accepted")
	}
}

func (o *object) RemoveValueForKey(key string) {
	delete(o.inner, key)
}

func (o *object) Keys() []string {
	return o.keys
}

func (o *object) MarshalEncode(output *jwriter.Writer) {
	output.RawByte('{')

	for index, key := range o.keys {
		if index > 0 {
			output.RawByte(',')
		}

		output.String(key)
		output.RawByte(':')

		o.NodeForKey(key).MarshalEncode(output)
	}

	output.RawByte('}')
}

func (o *object) MarshalWrite(output io.Writer) error {
	return MarshalWrite(o, output)
}
