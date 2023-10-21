package gosds

import (
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type placeholder struct {
	parent Node
	index  int
	root   Root
}

func newPlaceholder() *placeholder {
	return &placeholder{
		parent: nil,
		index:  0,
		root:   nil,
	}
}

func (p *placeholder) Set(val any) {
	if indexedParent, ok := p.Parent().AsIndexed(); ok {
		indexedParent.SetValueAtIndex(p.Index(), val)
	} else if root, ok := p.AsRoot(); ok {
		root.Set(val)
	}
}

func (p *placeholder) Parent() Node                  { return p.parent } //nolint:ireturn
func (p *placeholder) Index() int                    { return p.index }
func (p *placeholder) Get() any                      { return p }
func (p *placeholder) Primitive() any                { return nil }
func (p *placeholder) Exist() bool                   { return false }
func (p *placeholder) AsObject() (Object, bool)      { return nil, false }            //nolint:ireturn
func (p *placeholder) AsArray() (Array, bool)        { return nil, false }            //nolint:ireturn
func (p *placeholder) AsValue() (Value, bool)        { return nil, false }            //nolint:ireturn
func (p *placeholder) AsIndexed() (Indexed, bool)    { return nil, false }            //nolint:ireturn
func (p *placeholder) AsRoot() (Root, bool)          { return p.root, p.root != nil } //nolint:ireturn
func (p *placeholder) MustObject() Object            { return nil }                   //nolint:ireturn
func (p *placeholder) MustArray() Array              { return nil }                   //nolint:ireturn
func (p *placeholder) MustValue() Value              { return nil }                   //nolint:ireturn
func (p *placeholder) MustIndexed() Indexed          { return nil }                   //nolint:ireturn
func (p *placeholder) MustRoot() Root                { return p.root }                //nolint:ireturn
func (p *placeholder) MarshalEncode(*jwriter.Writer) { panic("") }
func (p *placeholder) MarshalWrite(io.Writer) error  { panic("") }
