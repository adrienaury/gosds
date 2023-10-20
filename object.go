package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Object interface {
	Node

	NodeForKey(key string) Node
	NodeAtIndex(index int) Node

	ValueForKey(key string) (any, bool)
	SetValueForKey(key string, value any)
	RemoveValueForKey(key string)

	ValueAtIndex(index int) (any, bool)
	SetValueAtIndex(index int, value any)
	RemoveValueAtIndex(index int)

	Keys() []string

	// PrimitiveObject returns a representation of the object as map[string]any
	PrimitiveObject() map[string]any
}

type object struct {
	values  []Node
	keys    []string
	indexes map[string]int

	parent Node
	index  int
}

func newObject() Object { //nolint:ireturn
	return &object{
		values:  []Node{},
		keys:    []string{},
		indexes: map[string]int{},
		parent:  nil,
		index:   0,
	}
}

func newObjectWithCapacity(capacity int) Object { //nolint:ireturn
	return &object{
		values:  make([]Node, 0, capacity),
		keys:    make([]string, 0, capacity),
		indexes: make(map[string]int, capacity),
		parent:  nil,
		index:   0,
	}
}

func (o *object) Parent() Node { //nolint:ireturn
	return o.parent
}

func (o *object) Index() int {
	return o.index
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

	for idx, key := range o.keys {
		result[key] = o.values[idx].Primitive()
	}

	return result
}

func (o *object) PrimitiveObject() map[string]any {
	return o.Primitive().(map[string]any) //nolint:forcetypeassert
}

func (o *object) NodeForKey(key string) Node { //nolint:ireturn
	return o.values[o.indexes[key]]
}

func (o *object) NodeAtIndex(index int) Node { //nolint:ireturn
	return o.values[index]
}

func (o *object) ValueForKey(key string) (any, bool) {
	idx, has := o.indexes[key]

	return o.values[idx].Value(), has
}

func (o *object) SetValueForKey(key string, val any) {
	index := len(o.keys)

	o.indexes[key] = index
	o.keys = append(o.keys, key)
	o.values = set(o.values, val, index, o)
}

func (o *object) RemoveValueForKey(key string) {
	if idx, has := o.indexes[key]; has {
		o.keys = append(o.keys[:idx], o.keys[idx+1:]...)
		o.values = append(o.values[:idx], o.values[idx+1:]...)
		delete(o.indexes, key)
	}
}

func (o *object) ValueAtIndex(index int) (any, bool) {
	if index >= len(o.values) {
		return nil, false
	}

	return o.values[index].Value(), true
}

func (o *object) SetValueAtIndex(index int, val any) {
	o.values = set(o.values, val, index, o)
}

func (o *object) RemoveValueAtIndex(index int) {
	if index >= len(o.values) {
		return
	}

	delete(o.indexes, o.keys[index])
	o.keys = append(o.keys[:index], o.keys[index+1:]...)
	o.values = append(o.values[:index], o.values[index+1:]...)
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

		o.values[index].MarshalEncode(output)
	}

	output.RawByte('}')
}

func (o *object) MarshalWrite(output io.Writer) error {
	return MarshalWrite(o, output)
}
