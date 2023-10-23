package gosds

type value struct {
	parent Node
	index  int
	key    string
	value  any
}

func newValue(val any) *value {
	return &value{
		parent: nil,
		index:  0,
		key:    "",
		value:  val,
	}
}

func (v *value) Root() Node { panic("") } //nolint:ireturn

func (v *value) Parent() Node { //nolint:ireturn
	return v.parent
}

func (v *value) Get() any                  { panic("") }
func (v *value) Set(val any)               { panic(val) }
func (v *value) Remove()                   { panic("") }
func (v *value) Primitive() any            { panic("") }
func (v *value) Exist() bool               { panic("") }
func (v *value) IsKeyed() bool             { panic("") }
func (v *value) IsIndexed() bool           { panic("") }
func (v *value) IsArray() bool             { panic("") }
func (v *value) IsObject() bool            { panic("") }
func (v *value) AsKeyed() Keyed            { panic("") } //nolint:ireturn
func (v *value) AsIndexed() Indexed        { panic("") } //nolint:ireturn
func (v *value) AsArray() Array            { panic("") } //nolint:ireturn
func (v *value) AsObject() Object          { panic("") } //nolint:ireturn
func (v *value) MarshalEncode(Encoder)     { panic("") }
func (v *value) MarshalWrite(Writer) error { panic("") }
