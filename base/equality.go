package base

import (
	"reflect"

	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

// CastToUint64 is a function that tries to cast value to uint64
// Casts any uint and any int, uintptr and any float
func CastToUint64(v any) Result[uint64] {
	intCase := func(v int64) Result[uint64] {
		if v < 0 {
			return Err[uint64](NewError("Cannot cast negative int to uint64"))
		}
		return Ok(uint64(v))
	}
	floatCase := func(v float64) Result[uint64] {
		if v > float64(GetMaxForType[uint64]()) {
			return Err[uint64](NewError("Cannot cast float greater than MaxUint64 to uint64"))
		}
		if v < float64(GetMinForType[uint64]()) {
			return Err[uint64](NewError("Cannot cast float smaller than MinUint64 to uint64"))
		}
		return Ok(uint64(v))
	}
	switch vt := v.(type) {
	case int:
		return intCase(int64(vt))
	case int8:
		return intCase(int64(vt))
	case int16:
		return intCase(int64(vt))
	case int32:
		return intCase(int64(vt))
	case int64:
		return intCase(int64(vt))
	case uint:
		return Ok(uint64(vt))
	case uint8:
		return Ok(uint64(vt))
	case uint16:
		return Ok(uint64(vt))
	case uint32:
		return Ok(uint64(vt))
	case uint64:
		return Ok(uint64(vt))
	case uintptr:
		return Ok(uint64(vt))
	case float32:
		return floatCase(float64(vt))
	case float64:
		return floatCase(float64(vt))
	default:
		return Err[uint64](NewError("Cannot cast to uint64"))
	}
}

// CastToInt64 is a function that tries to cast value to int64
// Casts any uint and any int, uintptr and any float
func CastToInt64(v any) Result[int64] {
	uintCase := func(v uint64) Result[int64] {
		if v > uint64(GetMaxForType[int64]()) {
			return Err[int64](NewError("Cannot cast uint greater than MaxInt64 to int64"))
		}
		return Ok(int64(v))
	}
	floatCase := func(v float64) Result[int64] {
		if v > float64(GetMaxForType[int64]()) {
			return Err[int64](NewError("Cannot cast float greater than MaxInt64 to int64"))
		}
		if v < float64(GetMinForType[int64]()) {
			return Err[int64](NewError("Cannot cast float smaller than MinInt64 to int64"))
		}
		return Ok(int64(v))
	}
	switch vt := v.(type) {
	case uint:
		return uintCase(uint64(vt))
	case uint8:
		return uintCase(uint64(vt))
	case uint16:
		return uintCase(uint64(vt))
	case uint32:
		return uintCase(uint64(vt))
	case uint64:
		return uintCase(uint64(vt))
	case uintptr:
		return uintCase(uint64(vt))
	case int:
		return Ok(int64(vt))
	case int8:
		return Ok(int64(vt))
	case int16:
		return Ok(int64(vt))
	case int32:
		return Ok(int64(vt))
	case int64:
		return Ok(int64(vt))
	case float32:
		return floatCase(float64(vt))
	case float64:
		return floatCase(float64(vt))
	default:
		return Err[int64](NewError("Cannot cast to int64"))
	}
}

// CastToFloat64 is a function that tries to cast value to int64
// Casts any uint and any int, uintptr and any float
func CastToFloat64(v any) Result[float64] {
	switch vt := v.(type) {
	case uint:
		return Ok(float64(vt))
	case uint8:
		return Ok(float64(vt))
	case uint16:
		return Ok(float64(vt))
	case uint32:
		return Ok(float64(vt))
	case uint64:
		return Ok(float64(vt))
	case uintptr:
		return Ok(float64(vt))
	case int:
		return Ok(float64(vt))
	case int8:
		return Ok(float64(vt))
	case int16:
		return Ok(float64(vt))
	case int32:
		return Ok(float64(vt))
	case int64:
		return Ok(float64(vt))
	case float32:
		return Ok(float64(vt))
	case float64:
		return Ok(float64(vt))
	default:
		return Err[float64](NewError("Cannot cast to float64"))
	}
}

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
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok(v1t == v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok(v1t == v2r.Unwrap())
		}
	case float32, float64:
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok(v1t == v2r.Unwrap())
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok(vv[0] == vv[1])
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok(vv[0] == vv[1])
		}
	case bool:
		if v2i, ok := v2.(bool); ok {
			return Ok(v1t == v2i)
		}
	case error:
		if v2i, ok := v2.(error); ok {
			return Ok(v1t == v2i)
		}
	case string:
		if v2i, ok := v2.(string); ok {
			return Ok(v1t == v2i)
		}
	}
	// Reflection based equality
	if reflect.DeepEqual(v1, v2) {
		return Ok(true)
	}
	// Couldn't be determined
	return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
}
