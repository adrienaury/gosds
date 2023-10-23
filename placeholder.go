package gosds

import (
	"io"
)

type placeholder struct {
	parent Node
	index  int
	key    string
	root   Root
}

func newPlaceholder() *placeholder {
	return &placeholder{
		parent: nil,
		index:  0,
		root:   nil,
		key:    "",
	}
}

func (p *placeholder) Set(val any) {
	switch {
	case p.Parent() != nil && p.Parent().IsKeyed():
		p.Parent().AsKeyed().SetValueForKey(p.Key(), val)
	case p.Parent() != nil && p.Parent().IsIndexed():
		p.Parent().AsIndexed().SetValueAtIndex(p.Index(), val)
	case p.IsRoot():
		p.AsRoot().Set(val)
	}
}

func (p *placeholder) Remove() {
	switch {
	case p.Parent() != nil && p.Parent().IsKeyed():
		p.Parent().AsKeyed().RemoveValueForKey(p.Key())
	case p.Parent() != nil && p.Parent().IsIndexed():
		p.Parent().AsIndexed().RemoveValueAtIndex(p.Index())
	case p.IsRoot():
		p.AsRoot().Remove()
	}
}

func (p *placeholder) Root() Node { //nolint:ireturn
	var result Node = p

	for result.Parent() != nil {
		result = result.Parent()
	}

	return result
}

func (p *placeholder) Parent() Node                 { return p.parent } //nolint:ireturn
func (p *placeholder) Index() int                   { return p.index }
func (p *placeholder) Key() string                  { return p.key }
func (p *placeholder) Get() any                     { return p }
func (p *placeholder) Primitive() any               { return nil }
func (p *placeholder) Exist() bool                  { return false }
func (p *placeholder) IsKeyed() bool                { return false }
func (p *placeholder) IsIndexed() bool              { return false }
func (p *placeholder) IsArray() bool                { return false }
func (p *placeholder) IsObject() bool               { return false }
func (p *placeholder) IsRoot() bool                 { return p.root != nil }
func (p *placeholder) AsKeyed() Keyed               { return nil }    //nolint:ireturn
func (p *placeholder) AsIndexed() Indexed           { return nil }    //nolint:ireturn
func (p *placeholder) AsArray() Array               { return nil }    //nolint:ireturn
func (p *placeholder) AsObject() Object             { return nil }    //nolint:ireturn
func (p *placeholder) AsRoot() Root                 { return p.root } //nolint:ireturn
func (p *placeholder) MarshalEncode(Encoder)        { panic("") }
func (p *placeholder) MarshalWrite(io.Writer) error { panic("") }
