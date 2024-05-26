package base

import "github.com/patrykjadamczyk/go-utils/datastructures/interfaces"

// EnsureComparableType check if value provided is comparable in some way
// Returns value and type of comparison needed
// 0 - not comparable
// 1 - default type -> comparable with ==
// 2 - nil
// 3 - obj.Equals(any(other))
// 4 - obj.EqualsAny(any(other))
func EnsureComparableType(value any) (any, uint8) {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return value, 1
	case uint, uint8, uint16, uint32, uint64:
		return value, 1
	case float32, float64:
		return value, 1
	case string:
		return value, 1
	case bool:
		return value, 1
	case uintptr:
		return value, 1
	case complex64, complex128:
		return value, 1
	case nil:
		return value, 2
	case interfaces.IEqualsAny:
		return value, 3
	case interfaces.IEqualsAnyFunc:
		return value, 4
	default:
		return value, 0
	}
}

// EqualityCheckUsingComparableType check if value provided is comparable in some way
// Returns tuple where first bool is condition and second bool is if it could be checked
func EqualityCheckUsingComparableType(value any, other any) (bool, bool) {
	switch v := value.(type) {
		// comparable types
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
		switch o := other.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return v == o, true
		case interfaces.IEqualsAny:
			return o.Equals(v), true
		case interfaces.IEqualsAnyFunc:
			return o.EqualsAny(v), true
		default:
			return false, false
		}
		// Value.Equals(other any)
	case interfaces.IEqualsAny:
		switch o := other.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return v.Equals(o), true
		case interfaces.IEqualsAny:
			return v.Equals(o), true
		case interfaces.IEqualsAnyFunc:
			return v.Equals(o), true
		default:
			return false, false
		}
		// Value.EqualsAny(other any)
	case interfaces.IEqualsAnyFunc:
		switch o := other.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool, uintptr, complex64, complex128, nil:
			return v.EqualsAny(o), true
		case interfaces.IEqualsAny:
			return v.EqualsAny(o), true
		case interfaces.IEqualsAnyFunc:
			return v.EqualsAny(o), true
		default:
			return false, false
		}
		// Couldn't be determined
	default:
		return false, false
	}
}
