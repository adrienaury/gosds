package gosds

type Array interface {
	Indexed
	AppendValue(value any)
	Node
}

type array struct {
	values []Node
	parent Node
	index  int
	key    string
	root   Root
}

func newArrayWithCapacity(capacity int) *array {
	return &array{
		values: make([]Node, 0, capacity),
		parent: nil,
		index:  0,
		key:    "",
		root:   nil,
	}
}

func (a *array) NodeAtIndex(index int) Node { //nolint:ireturn
	return a.values[index]
}

func (a *array) ValueAtIndex(index int) any {
	return a.values[index].Get()
}

func (a *array) SetValueAtIndex(index int, val any) {
	a.values = set(a.values, val, index, a)
}

func (a *array) RemoveValueAtIndex(index int) {
	a.values = append(a.values[:index], a.values[index+1:]...)
}

func (a *array) Size() int {
	return len(a.values)
}

func (a *array) PrimitiveArray() []any {
	return a.Primitive().([]any) //nolint:forcetypeassert
}

func (a *array) AppendValue(val any) {
	a.values = add(a.values, val, a)
}

func (a *array) Root() Node { //nolint:ireturn
	var result Node = a

	for result.Parent() != nil {
		result = result.Parent()
	}

	return result
}

func (a *array) Parent() Node { //nolint:ireturn
	return a.parent
}

func (a *array) Index() int {
	return a.index
}

func (a *array) Key() string {
	return a.key
}

func (a *array) Get() any {
	return a
}

func (a *array) Set(val any) {
	switch {
	case a.Parent() != nil && a.Parent().IsKeyed():
		a.Parent().AsKeyed().SetValueForKey(a.Key(), val)
	case a.Parent() != nil && a.Parent().IsIndexed():
		a.Parent().AsIndexed().SetValueAtIndex(a.Index(), val)
	case a.IsRoot():
		a.AsRoot().Set(val)
	}
}

func (a *array) Remove() {
	switch {
	case a.Parent() != nil && a.Parent().IsKeyed():
		a.Parent().AsKeyed().RemoveValueForKey(a.Key())
	case a.Parent() != nil && a.Parent().IsIndexed():
		a.Parent().AsIndexed().RemoveValueAtIndex(a.Index())
	case a.IsRoot():
		a.AsRoot().Remove()
	}
}

func (a *array) Primitive() any {
	result := make([]any, len(a.values))

	for index, val := range a.values {
		result[index] = val.Primitive()
	}

	return result
}

func (a *array) Exist() bool        { return true }
func (a *array) IsKeyed() bool      { return false }
func (a *array) IsIndexed() bool    { return true }
func (a *array) IsArray() bool      { return true }
func (a *array) IsObject() bool     { return false }
func (a *array) IsRoot() bool       { return a.root != nil }
func (a *array) AsKeyed() Keyed     { return nil }    //nolint:ireturn
func (a *array) AsIndexed() Indexed { return a }      //nolint:ireturn
func (a *array) AsArray() Array     { return a }      //nolint:ireturn
func (a *array) AsObject() Object   { return nil }    //nolint:ireturn
func (a *array) AsRoot() Root       { return a.root } //nolint:ireturn

func (a *array) MarshalEncode(output Encoder) {
	output.RawByte('[')

	for index, node := range a.values {
		if index > 0 {
			output.RawByte(',')
		}

		node.MarshalEncode(output)
	}

	output.RawByte(']')
}

func (a *array) MarshalWrite(output Writer) error {
	return MarshalWrite(a, output)
}
