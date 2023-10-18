package gosds

import (
	"fmt"
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Node interface {
	Parent() Node
	Value() any

	AsObject() (Object, bool)
	AsArray() (Array, bool)
	AsValue() (Value, bool)

	MustObject() Object
	MustArray() Array
	MustValue() Value

	// Primitive returns a representation of the node with following types :
	// - objects as map[string]any
	// - arrays as []any
	// - values as string, float64, bool or nil interface
	Primitive() any

	MarshalEncode(*jwriter.Writer)
	MarshalWrite(io.Writer) error
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
