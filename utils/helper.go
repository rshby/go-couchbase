package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

type Int interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func ExpectedNumber[T Int](v any) T {
	var result T
	switch value := v.(type) {
	case int:
		result = T(value)
	case int8:
		result = T(value)
	case int16:
		result = T(value)
	case int32:
		result = T(value)
	case int64:
		result = T(value)
	case uint:
		result = T(value)
	case uint8:
		result = T(value)
	case uint16:
		result = T(value)
	case uint32:
		result = T(value)
	case uint64:
		result = T(value)
	case uintptr:
		result = T(value)
	case float32:
		result = T(value)
	case float64:
		result = T(value)
	case string:
		result = T(StringToInt[T](value))
	default:
		result = 0
	}

	return T(result)
}

func ExpectString(v any) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		marshal, _ := json.Marshal(v)
		return string(marshal)
	}
}

// StringToInt converts a string to an integer value.
func StringToInt[T Int](s string) T {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return T(i)
}

func GenerateID() uint64 {
	return ExpectedNumber[uint64](time.Now().UnixNano() + int64(rand.Intn(10000)))
}
