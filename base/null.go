package base

import (
	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)


// UniversalNull is a function that determins current type and returns null value
//
// Returns Result type with:
// - Ok(v) type could be found
// - Ok(nil) if type doesn't have special case specified
// - ErrWithValue(nil, err) if this operation couldn't be determined
// This means that if you want you can use data value of result directly as null result
// And if you want to have special case for not determined operation you can check for res.IsError()
func UniversalNull(v any) Result[any] {
	// Shortcircuit for nil
	if v == nil {
		return ErrWithValue[any](nil, NewError("Value is nil"))
	}
	// Shortcircuit for Wrappers
	if vi, ok := v.(interfaces.IGetAny); ok {
		return UniversalNull(vi.Get())
	}
	if vi, ok := v.(interfaces.IGetAnyFunc); ok {
		return UniversalNull(vi.GetAny())
	}
	// Check for all basic go types
	// comparable types
	switch v.(type) {
	case int:
		return Ok[any](Null[int]().ValueOrZero())
	case int8:
		return Ok[any](Null[int8]().ValueOrZero())
	case int16:
		return Ok[any](Null[int16]().ValueOrZero())
	case int32:
		return Ok[any](Null[int32]().ValueOrZero())
	case int64:
		return Ok[any](Null[int64]().ValueOrZero())
	case uint:
		return Ok[any](Null[uint]().ValueOrZero())
	case uint8:
		return Ok[any](Null[uint8]().ValueOrZero())
	case uint16:
		return Ok[any](Null[uint16]().ValueOrZero())
	case uint32:
		return Ok[any](Null[uint32]().ValueOrZero())
	case uint64:
		return Ok[any](Null[uint64]().ValueOrZero())
	case float32:
		return Ok[any](Null[float32]().ValueOrZero())
	case float64:
		return Ok[any](Null[float64]().ValueOrZero())
	case complex64:
		return Ok[any](Null[complex64]().ValueOrZero())
	case complex128:
		return Ok[any](Null[complex128]().ValueOrZero())
	case bool:
		return Ok[any](Null[bool]().ValueOrZero())
	case string:
		return Ok[any](Null[string]().ValueOrZero())
	case error:
		return Ok[any](Null[error]().ValueOrZero())
	}
	return Ok[any](nil)
}
