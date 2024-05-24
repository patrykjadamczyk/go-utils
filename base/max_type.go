package base

import (
	"fmt"
	"math"
	"reflect"
)

type NumericType interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint
}

type SignedNumericType interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~int
}

type AnyFloat interface {
	~float32 | ~float64
}

type AnyInt interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint
}

type AnySint interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint
}

type AnyUint interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

// GetMaxForType
// Get Max Value for specified numeric type
func GetMaxForType[T NumericType]() T {
	// Note: This is wrapper arround implementation to trick generics in using just type parameter for getting this value
	var test T
	return EnsureType[T](getMaxForType(test))
}

// GetMaxForType Implementation
func getMaxForType(value any) any {
	switch value.(type) {
	case float64:
		return float64(math.MaxFloat64)
	case float32:
		return float32(math.MaxFloat32)
	case int:
		return int(math.MaxInt)
	case int8:
		return int8(math.MaxInt8)
	case int16:
		return int16(math.MaxInt16)
	case int32:
		return int32(math.MaxInt32)
	case int64:
		return int64(math.MaxInt64)
	case uint:
		return uint(math.MaxUint)
	case uint8:
		return uint8(math.MaxUint8)
	case uint16:
		return uint16(math.MaxUint16)
	case uint32:
		return uint32(math.MaxUint32)
	case uint64:
		return uint64(math.MaxUint64)
	default:
		panic(fmt.Sprintf("Invalid Type %s", reflect.TypeOf(value).Name()))
	}
}

// GetMinForType
// Get Min Value for specified numeric type
func GetMinForType[T NumericType]() T {
	// Note: This is wrapper arround implementation to trick generics in using just type parameter for getting this value
	var test T
	return EnsureType[T](getMinForType(test))
}

// GetMinForType Implementation
func getMinForType(value any) any {
	switch value.(type) {
	case float64:
		return -1 * float64(math.MaxFloat64)
	case float32:
		return -1 * float32(math.MaxFloat32)
	case int:
		return int(math.MinInt)
	case int8:
		return int8(math.MinInt8)
	case int16:
		return int16(math.MinInt16)
	case int32:
		return int32(math.MinInt32)
	case int64:
		return int64(math.MinInt64)
	case uint:
		return uint(0)
	case uint8:
		return uint8(0)
	case uint16:
		return uint16(0)
	case uint32:
		return uint32(0)
	case uint64:
		return uint64(0)
	default:
		panic(fmt.Sprintf("Invalid Type %s", reflect.TypeOf(value).Name()))
	}
}
