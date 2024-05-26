package base

import "github.com/patrykjadamczyk/go-utils/datastructures/interfaces"

// UniversalEquals is a function that determins if two values are equal
// It checks through == operator, different interfaces and other things
// Returns Result type with:
// - Ok(true) if values are equal
// - Ok(false) if values are not equal
// - ErrWithValue(false, err) if this operation couldn't be determined
// This means that if you want you can use data value of result directly as condition result
// And if you want to have special case for not determined operation you can check for res.IsError()
func UniversalEquals(v1 any, v2 any) Result[bool] {
	// Special shortcurcuit cases for v2
	switch v2t := v2.(type) {
		// Value.Get() any -> wrapped type
	case interfaces.IGetAnyFunc:
		return UniversalEquals(v1, v2t.GetAny())
		// Value.GetAny() any -> wrapped type
	case interfaces.IGetAny:
		return UniversalEquals(v1, v2t.Get())
	}
	switch v1t := v1.(type) {
		// comparable types - Go Types
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
		switch v2t := v2.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return Ok(v1t == v2t)
		case interfaces.IEqualsAny:
			return Ok(v2t.Equals(v1t))
		case interfaces.IEqualsAnyFunc:
			return Ok(v2t.EqualsAny(v1t))
		default:
			return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
		}
		// Value.Equals(other any)
	case interfaces.IEqualsAny:
		switch v2t := v2.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return Ok(v1t.Equals(v2t))
		case interfaces.IEqualsAny:
			return Ok(v1t.Equals(v2t))
		case interfaces.IEqualsAnyFunc:
			return Ok(v1t.Equals(v2t))
		default:
			return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
		}
		// Value.EqualsAny(other any)
	case interfaces.IEqualsAnyFunc:
		switch v2t := v2.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return Ok(v1t.EqualsAny(v2t))
		case interfaces.IEqualsAny:
			return Ok(v1t.EqualsAny(v2t))
		case interfaces.IEqualsAnyFunc:
			return Ok(v1t.EqualsAny(v2t))
		default:
			return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
		}
		// Value.Get() any -> wrapped type
	case interfaces.IGetAny:
		return UniversalEquals(v1t.Get(), v2)
		// Value.GetAny() any -> wrapped type
	case interfaces.IGetAnyFunc:
		return UniversalEquals(v1t.GetAny(), v2)
		// Couldn't be determined
	default:
		return ErrWithValue(false, NewError("Unsupported operation v1 == v2"))
	}
}
