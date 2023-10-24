package gosds

import "io"

type Value interface {
	Node
	Container
}

type value struct {
	parent Node
	index  int
	key    string
	value  any
	root   Root
}

func newValue(val any) *value {
	return &value{
		parent: nil,
		index:  0,
		key:    "",
		value:  val,
		root:   nil,
	}
}

func (v *value) Root() Root { //nolint:ireturn
	var result Node = v

	for result.Parent() != nil {
		result = result.Parent()
	}

	return result.AsRoot()
}

func (v *value) Parent() Node { //nolint:ireturn
	return v.parent
}

func (v *value) Index() int {
	return v.index
}

func (v *value) Key() string {
	return v.key
}

func (v *value) Get() any {
	return v.value
}

func (v *value) Set(val any) {
	switch val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		v.value = val
	default:
		switch {
		case v.Parent() != nil && v.Parent().IsKeyed():
			v.Parent().AsKeyed().SetValueForKey(v.Key(), val)
		case v.Parent() != nil && v.Parent().IsIndexed():
			v.Parent().AsIndexed().SetValueAtIndex(v.Index(), val)
		case v.IsRoot():
			v.AsRoot().Set(val)
		}
	}
}

func (v *value) Remove() {
	switch {
	case v.Parent() != nil && v.Parent().IsKeyed():
		v.Parent().AsKeyed().RemoveValueForKey(v.Key())
	case v.Parent() != nil && v.Parent().IsIndexed():
		v.Parent().AsIndexed().RemoveValueAtIndex(v.Index())
	case v.IsRoot():
		v.AsRoot().Remove()
	}
}

func (v *value) Primitive() any     { return v.value }
func (v *value) Exist() bool        { return true }
func (v *value) IsKeyed() bool      { return false }
func (v *value) IsIndexed() bool    { return false }
func (v *value) IsObject() bool     { return false }
func (v *value) IsArray() bool      { return false }
func (v *value) IsRoot() bool       { return v.root != nil }
func (v *value) AsKeyed() Keyed     { return nil }    //nolint:ireturn
func (v *value) AsIndexed() Indexed { return nil }    //nolint:ireturn
func (v *value) AsObject() Object   { return nil }    //nolint:ireturn
func (v *value) AsArray() Array     { return nil }    //nolint:ireturn
func (v *value) AsRoot() Root       { return v.root } //nolint:ireturn

func (v *value) MarshalEncode(output Encoder) { //nolint:cyclop
	switch typedValue := v.value.(type) {
	case string:
		output.String(typedValue)
	case int:
		output.Int(typedValue)
	case int64:
		output.Int64(typedValue)
	case int32:
		output.Int32(typedValue)
	case int16:
		output.Int16(typedValue)
	case int8:
		output.Int8(typedValue)
	case uint:
		output.Uint(typedValue)
	case uint64:
		output.Uint64(typedValue)
	case uint32:
		output.Uint32(typedValue)
	case uint16:
		output.Uint16(typedValue)
	case uint8:
		output.Uint8(typedValue)
	case float64:
		output.Float64(typedValue)
	case float32:
		output.Float32(typedValue)
	case bool:
		output.Bool(typedValue)
	case Number:
		output.RawString(typedValue.String())
	case nil:
		output.RawString("null")
	}
}

func (v *value) MarshalWrite(output io.Writer) error {
	return MarshalWrite(v, output)
}
