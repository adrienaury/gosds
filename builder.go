package gosds

import "errors"

var (
	ErrFinalized             = errors.New("node is finalized, can't be edited")
	ErrNoOpenedArrayOrObject = errors.New("can't close object or array because there is no opened node")
	ErrMissingKey            = errors.New("can't create value without key")
	ErrDoubleKey             = errors.New("already have a key for next value")
	ErrUnusedKey             = errors.New("key for value is unused")
)

type Builder struct {
	current   Node
	isObject  bool
	key       *string
	finalized Node
}

func NewBuilder() *Builder {
	return &Builder{
		current:   nil,
		isObject:  false,
		key:       nil,
		finalized: nil,
	}
}

func (b *Builder) AddKey(key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		return ErrNoOpenedArrayOrObject
	case b.isObject && b.key != nil:
		return ErrDoubleKey
	case b.isObject && b.key == nil:
		b.key = &key
	default:
		return ErrNoOpenedArrayOrObject
	}

	return nil
}

func (b *Builder) AddValue(val any) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = newValue(val)
	case b.isObject && b.key != nil:
		b.current.AsObject().SetValueForKey(*b.key, val)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.AsArray().AppendValue(val)
	}

	return nil
}

func (b *Builder) StartObject() error {
	return b.StartObjectWithCapacity(0)
}

func (b *Builder) StartObjectWithCapacity(capacity int) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = newObjectWithCapacity(capacity)
	case b.isObject && b.key != nil:
		object := newObjectWithCapacity(capacity)
		b.current.AsObject().SetValueForKey(*b.key, object)
		b.current = object
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		object := newObjectWithCapacity(capacity)
		b.current.AsArray().AppendValue(object)
		b.current = object
	}

	b.isObject = true

	return nil
}

func (b *Builder) StartArray() error {
	return b.StartArrayWithCapacity(0)
}

func (b *Builder) StartArrayWithCapacity(capacity int) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = newArrayWithCapacity(capacity)
	case b.isObject && b.key != nil:
		array := newArrayWithCapacity(capacity)
		b.current.AsObject().SetValueForKey(*b.key, array)
		b.current = array
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		array := newArrayWithCapacity(capacity)
		b.current.AsArray().AppendValue(array)
		b.current = array
	}

	b.isObject = false

	return nil
}

func (b *Builder) EndObjectOrArray() error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		return ErrNoOpenedArrayOrObject
	default:
		if b.current.Parent() == nil {
			b.finalized = b.current
		} else {
			b.current = b.current.Parent()
			b.isObject = b.current.AsObject() != nil
		}
	}

	return nil
}

func (b *Builder) Build() Root { //nolint:ireturn
	return newRoot(b.finalized)
}
