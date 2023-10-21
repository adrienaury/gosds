package gosds

import (
	"fmt"
	"io"

	"github.com/mailru/easyjson/jwriter"
)

type Node interface {
	Parented
	Container
	Castable
	JSONMarshaller
}

type Parented interface {
	Parent() Node
	Index() int
}

type Castable interface {
	AsObject() (Object, bool)
	AsArray() (Array, bool)
	AsValue() (Value, bool)
	AsIndexed() (Indexed, bool)
	AsContainer() (Container, bool)

	MustObject() Object
	MustArray() Array
	MustValue() Value
	MustIndexed() Indexed
	MustContainer() Container
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
