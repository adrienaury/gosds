package gosds

func NewValue(parent Node, val any) Node { //nolint:ireturn
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
