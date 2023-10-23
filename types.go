package gosds

import (
	"encoding/json"
	"io"
)

type (
	Writer = io.Writer
	Number = json.Number
)

type Accepted interface {
	~string | ~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32 | ~bool //nolint:lll
}
