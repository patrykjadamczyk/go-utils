package base

import (
	"reflect"

	compare "github.com/patrykjadamczyk/go-utils/base/CompareEnum"
	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

// UniversalLen is a function that returns length of any type
func UniversalCompare(v1 any, v2 any) compare.CompareEnum {
	// Shortcircuit for nil
	if v1 == nil && v2 == nil {
		return compare.Equal
	}
	if v1 == nil || v2 == nil {
		if v1 == nil {
			return compare.Less
		}
		if v2 == nil {
			return compare.Greater
		}
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return UniversalCompare(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return UniversalCompare(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return UniversalCompare(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return UniversalCompare(v1i.GetAny(), v2i.GetAny())
		}
	}
	// Shortcircuit for objects that implement comparision methods
	if v1i, ok := v1.(interfaces.IEqualsAny); ok {
		if v1i.Equals(v2) {
			return compare.Equal
		}
	}
	if v2i, ok := v2.(interfaces.IEqualsAny); ok {
		if v2i.Equals(v1) {
			return compare.Equal
		}
	}
	if v1i, ok := v1.(interfaces.IEqualsAnyFunc); ok {
		if v1i.EqualsAny(v2) {
			return compare.Equal
		}
	}
	if v2i, ok := v2.(interfaces.IEqualsAnyFunc); ok {
		if v2i.EqualsAny(v1) {
			return compare.Equal
		}
	}
	if v1i, ok := v1.(interfaces.ILessThanAny); ok {
		if v1i.LessThan(v2) {
			return compare.Less
		}
	}
	if v2i, ok := v2.(interfaces.ILessThanAny); ok {
		if v2i.LessThan(v1) {
			return compare.Greater
		}
	}
	if v1i, ok := v1.(interfaces.ILessThanAnyFunc); ok {
		if v1i.LessThanAny(v2) {
			return compare.Less
		}
	}
	if v2i, ok := v2.(interfaces.ILessThanAnyFunc); ok {
		if v2i.LessThanAny(v1) {
			return compare.Greater
		}
	}
	if v1i, ok := v1.(interfaces.IGreaterThanAny); ok {
		if v1i.GreaterThan(v2) {
			return compare.Greater
		}
	}
	if v2i, ok := v2.(interfaces.IGreaterThanAny); ok {
		if v2i.GreaterThan(v1) {
			return compare.Less
		}
	}
	if v1i, ok := v1.(interfaces.IGreaterThanAnyFunc); ok {
		if v1i.GreaterThanAny(v2) {
			return compare.Greater
		}
	}
	if v2i, ok := v2.(interfaces.IGreaterThanAnyFunc); ok {
		if v2i.GreaterThanAny(v1) {
			return compare.Less
		}
	}
	// Check for all basic go types
	// comparable types
	switch v1t := v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return compare.ToCompareEnum(v1r.Unwrap(), v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return compare.ToCompareEnum(v1r.Unwrap(), v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return compare.ToCompareEnum(v1r.Unwrap(), v2r.Unwrap())
		}
	case complex64, complex128:
		return compare.ComplexToCompareEnum(v1, v2)
	case bool:
		if v2i, ok := v2.(bool); ok {
			if v1t == v2i {
				return compare.Equal
			}
			return compare.Uncomparable
		}
	case error:
		if v2i, ok := v2.(error); ok {
			if v1t == v2i {
				return compare.Equal
			}
			return compare.Uncomparable
		}
	case string:
		if v2i, ok := v2.(string); ok {
			if v1t == v2i {
				return compare.Equal
			}
			return compare.Uncomparable
		}
	}
	// Reflection based equality
	if reflect.DeepEqual(v1, v2) {
		return compare.Equal
	}
	// Couldn't be determined
	return compare.Error
}
