package gosds

type Value interface {
	Node

	Value() any

	Set(val any)

	// Int64() (int64, bool)
	// MustInt64() int64
}

func newValue(parent Node, val any) Value { //nolint:ireturn
	return &value{
		parent: parent,
		value:  val,
	}
}

type value struct {
	parent Node
	value  any // can be string, float64, bool or nil interface
}

func (v *value) Parent() Node { //nolint:ireturn
	return v.parent
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

func (v *value) MustObject() Object { //nolint:ireturn
	return nil
}

func (v *value) MustArray() Array { //nolint:ireturn
	return nil
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
