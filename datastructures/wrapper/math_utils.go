package wrapper

import (
	"fmt"
	"strings"

	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
	"github.com/patrykjadamczyk/go-utils/math"
)

func add(v1, v2 any) Result[any] {
	// Shortcircuit for nil
	if v1 == nil && v2 == nil {
		return Ok[any](nil)
	}
	if v1 == nil {
		return Ok[any](v2)
	}
	if v2 == nil {
		return Ok[any](v1)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return add(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return add(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return add(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return add(v1i.GetAny(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IAddAny); ok {
		return Ok(v1i.Add(v2))
	}
	if v2i, ok := v2.(interfaces.IAddAny); ok {
		return Ok(v2i.Add(v1))
	}
	if v1i, ok := v1.(interfaces.IAddAnyFunc); ok {
		return Ok(v1i.AddAny(v2))
	}
	if v2i, ok := v2.(interfaces.IAddAnyFunc); ok {
		return Ok(v2i.AddAny(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1t := v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() + v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() + v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() + v2r.Unwrap())
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] + vv[1])
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] + vv[1])
		}
	case string:
		if v2i, ok := v2.(string); ok {
			return Ok[any](fmt.Sprintf("%s%s", v1t, v2i))
		}
	}
	// Couldn't be determined
	return ErrWithValue[any](nil, NewError("Unsupported operation v1 + v2"))
}

func substract(v1, v2 any) Result[any] {
	// Shortcircuit for nil
	if v1 == nil && v2 == nil {
		return Ok[any](nil)
	}
	if v1 == nil {
		return Ok[any](substract(0, v2))
	}
	if v2 == nil {
		return Ok[any](v1)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return substract(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return substract(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return substract(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return substract(v1i.GetAny(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.ISubstractAny); ok {
		return Ok(v1i.Substract(v2))
	}
	if v2i, ok := v2.(interfaces.ISubstractAny); ok {
		return Ok(v2i.Substract(v1))
	}
	if v1i, ok := v1.(interfaces.ISubstractAnyFunc); ok {
		return Ok(v1i.SubstractAny(v2))
	}
	if v2i, ok := v2.(interfaces.ISubstractAnyFunc); ok {
		return Ok(v2i.SubstractAny(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() - v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() - v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() - v2r.Unwrap())
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] - vv[1])
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] - vv[1])
		}
	}
	// Couldn't be determined
	return ErrWithValue[any](nil, NewError("Unsupported operation v1 - v2"))
}

func multiply(v1, v2 any) Result[any] {
	// Shortcircuit for nil
	if v1 == nil || v2 == nil {
		return Ok[any](nil)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return multiply(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return multiply(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return multiply(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return multiply(v1i.GetAny(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IMultiplyAny); ok {
		return Ok(v1i.Multiply(v2))
	}
	if v2i, ok := v2.(interfaces.IMultiplyAny); ok {
		return Ok(v2i.Multiply(v1))
	}
	if v1i, ok := v1.(interfaces.IMultiplyAnyFunc); ok {
		return Ok(v1i.MultiplyAny(v2))
	}
	if v2i, ok := v2.(interfaces.IMultiplyAnyFunc); ok {
		return Ok(v2i.MultiplyAny(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1t := v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() * v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() * v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() * v2r.Unwrap())
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] * vv[1])
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] * vv[1])
		}
	case string:
		switch v2t := v2.(type) {
		case int:
			return Ok[any](strings.Repeat(v1t, v2t))
		case int8:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case int16:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case int32:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case int64:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case uint:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case uint8:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case uint16:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case uint32:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case uint64:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case float32:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		case float64:
			return Ok[any](strings.Repeat(v1t, int(v2t)))
		}
	}
	// Couldn't be determined
	return ErrWithValue[any](nil, NewError("Unsupported operation v1 + v2"))
}

func divide(v1, v2 any) Result[any] {
	// Shortcircuit for nil
	if v2 == nil {
		return ErrWithValue[any](nil, NewError("Cannot divide by nil"))
	}
	if v1 == nil {
		return Ok[any](nil)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return divide(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return divide(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return divide(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return divide(v1i.GetAny(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IDivideAny); ok {
		return Ok(v1i.Divide(v2))
	}
	if v2i, ok := v2.(interfaces.IDivideAny); ok {
		return Ok(v2i.Divide(v1))
	}
	if v1i, ok := v1.(interfaces.IDivideAnyFunc); ok {
		return Ok(v1i.DivideAny(v2))
	}
	if v2i, ok := v2.(interfaces.IDivideAnyFunc); ok {
		return Ok(v2i.DivideAny(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() / v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() / v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() / v2r.Unwrap())
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] / vv[1])
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](vv[0] / vv[1])
		}
	}
	// Couldn't be determined
	return ErrWithValue[any](nil, NewError("Unsupported operation v1 / v2"))
}

func mod(v1, v2 any) Result[any] {
	// Shortcircuit for nil
	if v2 == nil {
		return ErrWithValue[any](nil, NewError("Cannot divide by nil"))
	}
	if v1 == nil {
		return Ok[any](nil)
	}
	// Shortcircuit for Wrappers
	if v1i, ok := v1.(interfaces.IGetAny); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return mod(v1i.Get(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return mod(v1i.Get(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IGetAnyFunc); ok {
		if v2i, ok := v2.(interfaces.IGetAny); ok {
			return mod(v1i.GetAny(), v2i.Get())
		}
		if v2i, ok := v2.(interfaces.IGetAnyFunc); ok {
			return mod(v1i.GetAny(), v2i.GetAny())
		}
	}
	if v1i, ok := v1.(interfaces.IModAny); ok {
		return Ok(v1i.Mod(v2))
	}
	if v2i, ok := v2.(interfaces.IModAny); ok {
		return Ok(v2i.Mod(v1))
	}
	if v1i, ok := v1.(interfaces.IModAnyFunc); ok {
		return Ok(v1i.ModAny(v2))
	}
	if v2i, ok := v2.(interfaces.IModAnyFunc); ok {
		return Ok(v2i.ModAny(v1))
	}
	// Check for all basic go types
	// comparable types
	switch v1t := v1.(type) {
	case int, int8, int16, int32, int64:
		v1r := CastToInt64(v1)
		v2r := CastToInt64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() % v2r.Unwrap())
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		v1r := CastToUint64(v1)
		v2r := CastToUint64(v2)
		if !v2r.IsError() {
			return Ok[any](v1r.Unwrap() % v2r.Unwrap())
		}
	case float32, float64:
		v1r := CastToFloat64(v1)
		v2r := CastToFloat64(v2)
		if !v2r.IsError() {
			return Ok[any](math.Fmod(v1r.Unwrap(), v2r.Unwrap()))
		}
	case complex64, complex128:
		if vv, ok := CastAll[complex64](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](
				complex64(complex(
					math.Fmod(real(vv[0]), real(vv[1])),
					math.Fmod(imag(vv[0]), imag(vv[1])),
				)),
			)
		}
		if vv, ok := CastAll[complex128](v1, v2); ok {
			Assert(len(vv) != 2)
			return Ok[any](
				complex(
					math.Fmod(real(vv[0]), real(vv[1])),
					math.Fmod(imag(vv[0]), imag(vv[1])),
				),
			)
		}
	case string:
		if v2i, ok := v2.(string); ok {
			return Ok[any](fmt.Sprintf(v1t, v2i))
		}
	}
	// Couldn't be determined
	return ErrWithValue[any](nil, NewError("Unsupported operation v1 % v2"))
}
