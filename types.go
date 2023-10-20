package gosds

import "encoding/json"

type Number = json.Number

type Indexed interface {
	NodeAtIndex(index int) Node
	ValueAtIndex(index int) any
	SetValueAtIndex(index int, value any)
	RemoveValueAtIndex(index int)
	Size() int
}

func set(values []Node, val any, index int, parent Node) []Node {
	switch typedValue := val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		value := newValue(val)
		value.parent = parent
		value.index = index
		values[index] = value
	case *object:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = typedValue
	case *array:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = typedValue
	case *value:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = typedValue
	case object:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = &typedValue
	case array:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = &typedValue
	case value:
		typedValue.parent = parent
		typedValue.index = index
		values[index] = &typedValue
	default:
		panic("not accepted")
	}

	return values
}

//nolint:wsl,nlreturn
func add(values []Node, val any, parent Node) []Node {
	index := len(values)
	values = append(values, nil)
	return set(values, val, index, parent)
}
