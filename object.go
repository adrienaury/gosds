package gosds

import (
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Object interface {
	Keyed
	Indexed
	Node
}

type object struct {
	list   *orderedmap.OrderedMap[string, Node]
	parent Node
	index  int
	key    string
	root   Root
}

func newObjectWithCapacity(capacity int) *object {
	return &object{
		list:   orderedmap.New[string, Node](capacity),
		parent: nil,
		index:  0,
		key:    "",
		root:   nil,
	}
}

func (o *object) NodeForKey(key string) Node { //nolint:ireturn
	val, has := o.list.Get(key)
	if !has {
		placeholder := newPlaceholder()
		o.SetValueForKey(key, placeholder)

		return placeholder
	}

	return val
}

func (o *object) ValueForKey(key string) (any, bool) {
	node := o.NodeForKey(key)

	if node == nil {
		return nil, false
	}

	return node.Get(), true
}

func (o *object) SetValueForKey(key string, val any) {
	o.list.Set(key, accept(val, key, 0, o))
}

func (o *object) RemoveValueForKey(key string) {
	o.list.Delete(key)
}

func (o *object) Keys() []string {
	result := make([]string, 0, o.Size())

	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		result = append(result, pair.Key)
	}

	return result
}

func (o *object) PrimitiveObject() map[string]any {
	return o.Primitive().(map[string]any) //nolint:forcetypeassert
}

func (o *object) NodeAtIndex(index int) Node { //nolint:ireturn
	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		if pair.Value.Index() == index {
			return pair.Value
		}
	}

	return nil
}

func (o *object) ValueAtIndex(index int) any {
	node := o.NodeAtIndex(index)

	if node != nil {
		return node.Get()
	}

	return nil
}

func (o *object) SetValueAtIndex(index int, value any) {
	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		if pair.Value.Index() == index {
			o.SetValueForKey(pair.Key, value)
		}
	}
}

func (o *object) RemoveValueAtIndex(index int) {
	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		if pair.Value.Index() == index {
			o.list.Delete(pair.Key)
		}
	}
}

func (o *object) Size() int {
	return o.list.Len()
}

func (o *object) PrimitiveArray() []any {
	result := make([]any, 0, o.Size())

	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		result = append(result, pair.Value.Primitive())
	}

	return result
}

func (o *object) Root() Node { //nolint:ireturn
	var result Node = o

	for result.Parent() != nil {
		result = result.Parent()
	}

	return result
}

func (o *object) Parent() Node { //nolint:ireturn
	return o.parent
}

func (o *object) Index() int {
	return o.index
}

func (o *object) Key() string {
	return o.key
}

func (o *object) Get() any {
	return o
}

func (o *object) Set(val any) {
	switch {
	case o.Parent() != nil && o.Parent().IsKeyed():
		o.Parent().AsKeyed().SetValueForKey(o.Key(), val)
	case o.Parent() != nil && o.Parent().IsIndexed():
		o.Parent().AsIndexed().SetValueAtIndex(o.Index(), val)
	case o.IsRoot():
		o.AsRoot().Set(val)
	}
}

func (o *object) Remove() {
	switch {
	case o.Parent() != nil && o.Parent().IsKeyed():
		o.Parent().AsKeyed().RemoveValueForKey(o.Key())
	case o.Parent() != nil && o.Parent().IsIndexed():
		o.Parent().AsIndexed().RemoveValueAtIndex(o.Index())
	case o.IsRoot():
		o.AsRoot().Remove()
	}
}

func (o *object) Primitive() any {
	result := make(map[string]any, o.Size())

	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		result[pair.Key] = pair.Value.Primitive()
	}

	return result
}

func (o *object) Exist() bool        { return true }
func (o *object) IsKeyed() bool      { return true }
func (o *object) IsIndexed() bool    { return true }
func (o *object) IsArray() bool      { return false }
func (o *object) IsObject() bool     { return true }
func (o *object) IsRoot() bool       { return o.root != nil }
func (o *object) AsKeyed() Keyed     { return o }      //nolint:ireturn
func (o *object) AsIndexed() Indexed { return o }      //nolint:ireturn
func (o *object) AsArray() Array     { return nil }    //nolint:ireturn
func (o *object) AsObject() Object   { return o }      //nolint:ireturn
func (o *object) AsRoot() Root       { return o.root } //nolint:ireturn

func (o *object) MarshalEncode(output Encoder) {
	output.RawByte('{')

	printComma := false

	for pair := o.list.Oldest(); pair != nil; pair = pair.Next() {
		if _, ok := pair.Value.(*placeholder); ok {
			continue
		}

		if printComma {
			output.RawByte(',')
		} else {
			printComma = true
		}

		output.String(pair.Key)
		output.RawByte(':')

		pair.Value.MarshalEncode(output)
	}

	output.RawByte('}')
}

func (o *object) MarshalWrite(output Writer) error {
	return MarshalWrite(o, output)
}
