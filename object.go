package gosds

import "encoding/json"

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

func NewObject(parent Node) Object { //nolint:ireturn
	return &object{
		inner:  map[string]Node{},
		keys:   []string{},
		parent: parent,
	}
}

func NewObjectWithCapacity(parent Node, capacity int) Object { //nolint:ireturn
	return &object{
		inner:  make(map[string]Node, capacity),
		keys:   make([]string, 0, capacity),
		parent: parent,
	}
}

type object struct {
	inner map[string]Node
	keys  []string

	parent Node
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

func (o *object) MustObject() Object { //nolint:ireturn
	return o
}

func (o *object) MustArray() Array { //nolint:ireturn
	return nil
}

func (o *object) Primitive() any {
	result := make(map[string]any, len(o.keys))

	for key, val := range o.inner {
		switch typedVal := val.(type) {
		case Node:
			result[key] = typedVal.Primitive()
		default:
			result[key] = val
		}
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
	switch typedValue := value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, json.Number:
		o.inner[key] = NewValue(o, value)
	case Node:
		o.inner[key] = typedValue
	default:
		panic("unimplemented")
	}
}

func (o *object) RemoveValueForKey(key string) {
	delete(o.inner, key)
}

func (o *object) Keys() []string {
	return o.keys
}
