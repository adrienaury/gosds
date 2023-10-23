package gosds

import (
	"fmt"
	"io"

	"github.com/bytedance/sonic/ast"
	"github.com/mailru/easyjson/jwriter"
)

type Marshaler interface {
	MarshalEncode(Encoder)
	MarshalWrite(Writer) error
}

type Encoder interface {
	EncoderRaw
	EncoderString
	EncoderInt
	EncoderUint
	EncoderFloat
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
	Uint32(n uint32)
	Uint16(n uint16)
	Uint8(n uint8)
}

type EncoderFloat interface {
	Float64(n float64)
	Float32(n float32)
}

type EncoderBool interface {
	Bool(v bool)
}

func MarshalWrite(node Node, output io.Writer) error {
	writer := &jwriter.Writer{NoEscapeHTML: true} //nolint:exhaustruct
	node.MarshalEncode(writer)

	if writer.Error != nil {
		return fmt.Errorf("%w", writer.Error)
	}

	if _, err := writer.DumpTo(output); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func Unmarshal(jstr string) (Root, error) {
	builder := &SonicBuilder{Builder: Builder{}}

	ast.Preorder(jstr, builder, &ast.VisitorOptions{OnlyNumber: true})

	return builder.Build(), nil
}
