package gosds

import "encoding/json"

type SonicBuilder struct {
	Builder
}

func NewBuilderSonic() *SonicBuilder {
	return &SonicBuilder{
		Builder: *NewBuilder(),
	}
}

func (b *SonicBuilder) OnNull() error {
	return b.AddValue(nil)
}

func (b *SonicBuilder) OnBool(v bool) error {
	return b.AddValue(v)
}

func (b *SonicBuilder) OnString(v string) error {
	return b.AddValue(v)
}

func (b *SonicBuilder) OnInt64(_ int64, n json.Number) error {
	return b.AddValue(n)
}

func (b *SonicBuilder) OnFloat64(_ float64, n json.Number) error {
	return b.AddValue(n)
}

func (b *SonicBuilder) OnObjectBegin(capacity int) error {
	return b.StartObjectWithCapacity(capacity)
}

func (b *SonicBuilder) OnObjectKey(key string) error {
	return b.AddKey(key)
}

func (b *SonicBuilder) OnObjectEnd() error {
	return b.EndObjectOrArray()
}

func (b *SonicBuilder) OnArrayBegin(capacity int) error {
	return b.StartArrayWithCapacity(capacity)
}

func (b *SonicBuilder) OnArrayEnd() error {
	return b.EndObjectOrArray()
}
