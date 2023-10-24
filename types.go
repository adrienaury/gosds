package gosds

import (
	"encoding/json"
	"fmt"
	"io"
)

const defaultCapacity = 5

type (
	Writer = io.Writer
	Number = json.Number
)

type Accepted interface {
	~string | ~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32 | ~bool //nolint:lll
}

func accept(val any, key string, index int, parent Node) Node { //nolint:ireturn
	switch typedValue := val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		value := newValue(val)
		value.parent = parent
		value.index = index
		value.key = key

		return value
	case *object:
		typedValue.parent = parent
		typedValue.index = index
		typedValue.key = key

		return typedValue
	case *array:
		typedValue.parent = parent
		typedValue.index = index
		typedValue.key = key

		return typedValue
	case *value:
		typedValue.parent = parent
		typedValue.index = index
		typedValue.key = key

		return typedValue
	case *placeholder:
		typedValue.parent = parent
		typedValue.index = index
		typedValue.key = key

		return typedValue
	default:
		panic(fmt.Sprintf("type not accepted : %T", val))
	}
}

func set(values []Node, val any, index int, parent Node) []Node {
	values[index] = accept(val, "", index, parent)

	return values
}

func add(values []Node, val any, parent Node) []Node {
	index := len(values)
	values = append(values, nil)

	return set(values, val, index, parent)
}
