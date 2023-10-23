package gosds

type Marshaler interface {
	MarshalEncode(Encoder)
	MarshalWrite(Writer) error
}

type Encoder interface {
	EncoderRaw
	EncoderString
	EncoderInt
	EncoderUint
	EncoderBool
}

type EncoderRaw interface {
	RawByte(c byte)
	RawString(s string)
}

type EncoderString interface {
	String(s string)
}

type EncoderInt interface {
	Int(n int)
	Int64(n int64)
	Int32(n int32)
	Int16(n int16)
	Int8(n int8)
}

type EncoderUint interface {
	Uint(n uint)
	Uint64(n uint64)
	Uint32Str(n uint32)
	Uint16Str(n uint16)
	Uint8Str(n uint8)
}

type EncoderFloat interface {
	Float64(n float64)
	Float32(n float32)
}

type EncoderBool interface {
	Bool(v bool)
}
