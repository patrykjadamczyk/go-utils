package base

import (
	"reflect"

	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

// UniversalLen is a function that returns length of any type
func UniversalLen(v any) int {
	// Shortcircuit for nil
	if v == nil {
		return 0
	}
	// Shortcircuit for Wrappers
	if vi, ok := v.(interfaces.IGetAny); ok {
		return UniversalLen(vi.Get())
	}
	if vi, ok := v.(interfaces.IGetAnyFunc); ok {
		return UniversalLen(vi.GetAny())
	}
	// Shortcircuit for objects that implement Len() int or Len() uint
	if vi, ok := v.(interfaces.ILenInt); ok {
		return vi.Len()
	}
	if vi, ok := v.(interfaces.ILenUint); ok {
		l := vi.Len()
		AssertCustomError(l > uint(GetMaxForType[int]()), NewError("Cannot get len of uint greater than MaxInt"))
		return int(vi.Len())
	}
	// Test for basic types
	switch vt := v.(type) {
	case []int:
		return len(vt)
	case []int8:
		return len(vt)
	case []int16:
		return len(vt)
	case []int32:
		return len(vt)
	case []int64:
		return len(vt)
	case []uint:
		return len(vt)
	case []uint8:
		return len(vt)
	case []uint16:
		return len(vt)
	case []uint32:
		return len(vt)
	case []uint64:
		return len(vt)
	case []float32:
		return len(vt)
	case []float64:
		return len(vt)
	case []complex64:
		return len(vt)
	case []complex128:
		return len(vt)
	case []uintptr:
		return len(vt)
	case []string:
		return len(vt)
	case []bool:
		return len(vt)
	case []any:
		return len(vt)
	}
	// If not found, fall back on reflection
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len()
	}
	return 0
}
