package gosds

import "encoding/json"

type Number = json.Number

//nolint:wsl,nlreturn
func set(values []Node, val any, index int, parent Node) []Node {
	switch typedValue := val.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, Number, nil:
		value := newValue(val)
		value.parent = parent
		value.index = index
		return append(values, value)
	case *object:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, typedValue)
	case *array:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, typedValue)
	case *value:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, typedValue)
	case object:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, &typedValue)
	case array:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, &typedValue)
	case value:
		typedValue.parent = parent
		typedValue.index = index
		return append(values, &typedValue)
	default:
		panic("not accepted")
	}
}
