package gosds

import (
	"fmt"
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Node interface {
	Parent() Node
	Index() int
	Value() any

	// Primitive returns a representation of the node with following types :
	// - objects as map[string]any
	// - arrays as []any
	// - values as any
	Primitive() any

	Castable

	JSONMarshaller
}

type Castable interface {
	AsObject() (Object, bool)
	AsArray() (Array, bool)
	AsValue() (Value, bool)
	AsIndexed() (Indexed, bool)

	MustObject() Object
	MustArray() Array
	MustValue() Value
	MustIndexed() Indexed
}

type JSONMarshaller interface {
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
