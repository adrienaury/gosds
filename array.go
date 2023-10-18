package gosds

type Array interface {
	Node

	ValueAtIndex(index int) any
	SetValueAtIndex(index int, value any)
	AppendValue(value any)

	Size() int

	// PrimitiveArray returns a representation of the array as []any
	PrimitiveArray() []any
}

func NewArray(parent Node) Array { //nolint:ireturn
	return &array{
		values: []any{},
		parent: parent,
	}
}

func NewArrayWithCapacity(parent Node, capacity int) Array { //nolint:ireturn
	return &array{
		values: make([]any, capacity),
		parent: parent,
	}
}

type array struct {
	values []any

	parent Node
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

func (a *array) MustObject() Object { //nolint:ireturn
	return nil
}

func (a *array) MustArray() Array { //nolint:ireturn
	return a
}

func (a *array) Primitive() any {
	result := make([]any, len(a.values))

	for index, val := range a.values {
		switch typedVal := val.(type) {
		case Node:
			result[index] = typedVal.Primitive()
		default:
			result[index] = val
		}
	}

	return result
}

func (a *array) PrimitiveArray() []any {
	return a.Primitive().([]any) //nolint:forcetypeassert
}

func (a *array) ValueAtIndex(index int) any {
	return a.values[index]
}

func (a *array) SetValueAtIndex(index int, value any) {
	switch value.(type) {
	case string, int, bool:
		a.values[index] = NewValue(a, value)
	default:
		panic("unimplemented")
	}
}

func (a *array) AppendValue(value any) {
	switch value.(type) {
	case string, int, bool:
		a.values = append(a.values, NewValue(a, value))
	default:
		panic("unimplemented")
	}
}

func (a *array) Size() int {
	return len(a.values)
}
