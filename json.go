package gosds

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

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

type Decoder interface {
	More() bool
	Token() (json.Token, error)
}

func Marshal(node Node) (string, error) {
	builder := &strings.Builder{}

	if err := MarshalWrite(node, builder); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func MarshalWrite(node Node, output Writer) error {
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

func MarshalEncode(node Node, output Encoder) {
	node.MarshalEncode(output)
}

func Unmarshal(jstr string) (Root, error) { //nolint:ireturn
	builder := NewBuilderSonic()

	if err := ast.Preorder(jstr, builder, &ast.VisitorOptions{OnlyNumber: true}); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return builder.Build(), nil
}

func UnmarshalRead(input Reader) (Root, error) { //nolint:ireturn
	scanner := bufio.NewScanner(input)

	if scanner.Scan() {
		return Unmarshal(scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("%w", scanner.Err())
	}

	return nil, io.EOF
}

func UnmarshalDecode(_ Decoder) error {
	panic("not implemented")
}
