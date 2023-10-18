package gosds

import (
	"encoding/json"
	"errors"
)

type Builder interface {
	AddNull(key string) error
	AddBool(v bool, key string) error
	AddString(v string, key string) error
	AddInt(v int64, key string) error
	AddFloat(v float64, key string) error
	AddNumber(n json.Number, key string) error
	StartObject(key string) error
	StartArray(key string) error

	EndObjectOrArray() error

	Finalize() Node
}

type SonicBuilder interface {
	OnNull() error
	OnBool(v bool) error
	OnString(v string) error
	OnInt64(v int64, n json.Number) error
	OnFloat64(v float64, n json.Number) error
	OnObjectBegin(capacity int) error
	OnObjectKey(key string) error
	OnObjectEnd() error
	OnArrayBegin(capacity int) error
	OnArrayEnd() error
}

var (
	ErrFinalized             = errors.New("node is finalized, can't be edited")
	ErrNoOpenedArrayOrObject = errors.New("can't close object or array because there is no opened node")
	ErrMissingKey            = errors.New("can't create value without key")
	ErrDoubleKey             = errors.New("already have a key for value")
	ErrUnusedKey             = errors.New("key for value is unused")
)

type builder struct {
	current   Node
	isObject  bool // true = object, false = array
	finalized Node

	// for SonicBuilder
	key *string
}

func NewBuilder() Builder { //nolint:ireturn
	return &builder{
		current:   nil,
		isObject:  false,
		finalized: nil,
		key:       nil,
	}
}

func NewBuilderSonic() SonicBuilder { //nolint:ireturn
	return &builder{
		current:   nil,
		isObject:  false,
		finalized: nil,
		key:       nil,
	}
}

func (b *builder) AddNull(key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, nil)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, nil)
	default:
		b.current.MustArray().AppendValue(nil)
	}

	return nil
}

func (b *builder) AddBool(val bool, key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, val)
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) AddFloat(val float64, key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, val)
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) AddInt(val int64, key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, val)
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) AddNumber(num json.Number, key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, num)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, num)
	default:
		b.current.MustArray().AppendValue(num)
	}

	return nil
}

func (b *builder) AddString(val string, key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject:
		b.current.MustObject().SetValueForKey(key, val)
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) EndObjectOrArray() error {
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
			_, b.isObject = b.current.AsObject()
		}
	}

	return nil
}

func (b *builder) StartArray(key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = NewArray(nil)
	case b.isObject:
		array := NewArray(b.current)
		b.current.MustObject().SetValueForKey(key, array)
		b.current = array
	default:
		array := NewArray(b.current)
		b.current.MustArray().AppendValue(array)
		b.current = array
	}

	b.isObject = false

	return nil
}

func (b *builder) StartObject(key string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = NewObject(nil)
	case b.isObject:
		object := NewObject(b.current)
		b.current.MustObject().SetValueForKey(key, object)
		b.current = object
	default:
		object := NewObject(b.current)
		b.current.MustArray().AppendValue(object)
		b.current = object
	}

	b.isObject = true

	return nil
}

func (b *builder) Finalize() Node { //nolint:ireturn
	return b.finalized
}

func (b *builder) OnNull() error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, nil)
	case b.isObject && b.key != nil:
		b.current.MustObject().SetValueForKey(*b.key, nil)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.MustArray().AppendValue(nil)
	}

	return nil
}

func (b *builder) OnBool(val bool) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject && b.key != nil:
		b.current.MustObject().SetValueForKey(*b.key, val)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) OnString(val string) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, val)
	case b.isObject && b.key != nil:
		b.current.MustObject().SetValueForKey(*b.key, val)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.MustArray().AppendValue(val)
	}

	return nil
}

func (b *builder) OnInt64(_ int64, num json.Number) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, num)
	case b.isObject && b.key != nil:
		b.current.MustObject().SetValueForKey(*b.key, num)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.MustArray().AppendValue(num)
	}

	return nil
}

func (b *builder) OnFloat64(_ float64, num json.Number) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.finalized = NewValue(nil, num)
	case b.isObject && b.key != nil:
		b.current.MustObject().SetValueForKey(*b.key, num)
		b.key = nil
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		b.current.MustArray().AppendValue(num)
	}

	return nil
}

func (b *builder) OnObjectBegin(capacity int) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = NewObjectWithCapacity(nil, capacity)
	case b.isObject && b.key != nil:
		object := NewObjectWithCapacity(b.current, capacity)
		b.current.MustObject().SetValueForKey(*b.key, object)
		b.current = object
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		object := NewObjectWithCapacity(b.current, capacity)
		b.current.MustArray().AppendValue(object)
		b.current = object
	}

	b.isObject = true

	return nil
}

func (b *builder) OnObjectKey(key string) error {
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

func (b *builder) OnObjectEnd() error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		return ErrNoOpenedArrayOrObject
	case b.isObject && b.key != nil:
		return ErrUnusedKey
	case b.isObject && b.key == nil:
		if b.current.Parent() == nil {
			b.finalized = b.current
		} else {
			b.current = b.current.Parent()
			_, b.isObject = b.current.AsObject()
		}
	default:
		return ErrNoOpenedArrayOrObject
	}

	return nil
}

func (b *builder) OnArrayBegin(capacity int) error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		b.current = NewArrayWithCapacity(nil, capacity)
	case b.isObject && b.key != nil:
		array := NewArrayWithCapacity(b.current, capacity)
		b.current.MustObject().SetValueForKey(*b.key, array)
		b.current = array
	case b.isObject && b.key == nil:
		return ErrMissingKey
	default:
		array := NewArrayWithCapacity(b.current, capacity)
		b.current.MustArray().AppendValue(array)
		b.current = array
	}

	b.isObject = false

	return nil
}

func (b *builder) OnArrayEnd() error {
	switch {
	case b.finalized != nil:
		return ErrFinalized
	case b.current == nil:
		return ErrNoOpenedArrayOrObject
	case b.isObject:
		return ErrNoOpenedArrayOrObject
	default:
		if b.current.Parent() == nil {
			b.finalized = b.current
		} else {
			b.current = b.current.Parent()
			_, b.isObject = b.current.AsObject()
		}
	}

	return nil
}
