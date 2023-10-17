package gosds

func NewValue(parent Node, val any) Node { //nolint:ireturn
	return &value{
		parent: parent,
		value:  val,
	}
}

type value struct {
	parent Node
	value  any
}

func (v *value) Parent() Node { //nolint:ireturn
	return v.parent
}

func (v *value) Value() any {
	return v.value
}
