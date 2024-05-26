package base

import (
	"reflect"

	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

// UniversalEquals is a function that determins if two values are equal
// It checks through == operator, different interfaces and other things
// Returns Result type with:
// - Ok(true) if values are equal
// - Ok(false) if values are not equal
// - ErrWithValue(false, err) if this operation couldn't be determined
// This means that if you want you can use data value of result directly as condition result
// And if you want to have special case for not determined operation you can check for res.IsError()
func UniversalEquals(v1 any, v2 any) Result[bool] {
	// Shortcircuit for nil
	if v1 == nil || v2 == nil {
		return Ok(v1 == nil && v2 == nil)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return UniversalEquals(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return UniversalEquals(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return UniversalEquals(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return UniversalEquals(v1i.GetAny(), v2i.GetAny())
		}
	}
	// Shortcircuit for objects that implement Equals(any) or EqualsAny(any)
	if v1i, ok := v1.(interfaces.IEqualsAny); ok {
		return Ok(v1i.Equals(v2))
	}
	if v2i, ok := v2.(interfaces.IEqualsAny); ok {
		return Ok(v2i.Equals(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1t := v1.(type) {
	case int, int8, int16, int32, int64:
		switch v2t := v2.(type) {
		case int, int8, int16, int32, int64:
			return Ok(v1t == v2t)
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		switch v2t := v2.(type) {
		case uint, uint8, uint16, uint32, uint64, uintptr:
			return Ok(v1t == v2t)
		}

	case float32, float64, string, bool, complex64, complex128, error:
		switch v2t := v2.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, error:
			return Ok(v1t == v2t)
		}
	}
	// Reflection based equality
	if reflect.DeepEqual(v1, v2) {
		return Ok(true)
	}
	// Couldn't be determined
	return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
}
