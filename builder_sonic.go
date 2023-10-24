package gosds

import "encoding/json"

type BuilderSonic interface { //nolint:interfacebloat
	OnNull() error
	OnBool(v bool) error
	OnString(v string) error
	OnInt64(_ int64, n json.Number) error
	OnFloat64(_ float64, n json.Number) error
	OnObjectBegin(capacity int) error
	OnObjectKey(key string) error
	OnObjectEnd() error
	OnArrayBegin(capacity int) error
	OnArrayEnd() error
	Build() Root
}

func NewBuilderSonic() BuilderSonic { //nolint:ireturn
	return &builder{
		current:   nil,
		isObject:  false,
		key:       nil,
		finalized: nil,
	}
}

func (b *builder) OnNull() error {
	return b.AddValue(nil)
}

func (b *builder) OnBool(v bool) error {
	return b.AddValue(v)
}

func (b *builder) OnString(v string) error {
	return b.AddValue(v)
}

func (b *builder) OnInt64(_ int64, n json.Number) error {
	return b.AddValue(n)
}

func (b *builder) OnFloat64(_ float64, n json.Number) error {
	return b.AddValue(n)
}

func (b *builder) OnObjectBegin(capacity int) error {
	return b.StartObjectWithCapacity(capacity)
}

func (b *builder) OnObjectKey(key string) error {
	return b.AddKey(key)
}

func (b *builder) OnObjectEnd() error {
	return b.EndObjectOrArray()
}

func (b *builder) OnArrayBegin(capacity int) error {
	return b.StartArrayWithCapacity(capacity)
}

func (b *builder) OnArrayEnd() error {
	return b.EndObjectOrArray()
}
