package gosds

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

func NewObject() Object { //nolint:ireturn
	return &object{
		inner:  map[string]Node{},
		keys:   []string{},
		parent: nil,
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

func (o *object) Primitive() any {
	panic("not implemented")
}

func (o *object) PrimitiveObject() map[string]any {
	panic("not implemented")
}

func (o *object) NodeForKey(key string) Node { //nolint:ireturn
	return o.inner[key]
}

func (o *object) ValueForKey(key string) (any, bool) {
	node, has := o.inner[key]

	return node.Value(), has
}

func (o *object) SetValueForKey(key string, value any) {
	switch value.(type) {
	case string, int, bool:
		o.inner[key] = NewValue(o, value)
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
