package base

// Any function
type AnyFunc = func(any) any

// Any function returning AnyResult
type AnyFuncResult = func(any) AnyResult

// Any function returning AnyNullable
type AnyFuncNullable = func(any) AnyNullable

// Any action with arguments
type AnyAction = func(any)

// Action function
type Action = func()

// Any Result
type AnyResult = Result[any]

// Any Nullable
type AnyNullable = Nullable[any]

// Rust like types

type F32 float32
type F64 float64
type I8 int8
type I16 int16
type I32 int32
type I64 int64
type Str string
type U8 uint8
type U16 uint16
type U32 uint32
type U64 uint64

// Capitalized Types

type Float32 float32
type Float64 float64
type Float float64

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Int int

type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Uint uint

type Sint8 int8
type Sint16 int16
type Sint32 int32
type Sint64 int64
type Sint int


// Ensure that value is of type T or panic
func EnsureType[T any](value any) T {
	switch value := value.(type) {
	case T:
		return value
	default:
		panic("Invalid Type")
	}
}

// Ensure that value is of type T and return Result
func EnsureTypeResult[T any](value any) Result[T] {
	if val, ok := value.(T); ok {
		return Ok(val)
	}
	return Err[T](NewError("Invalid Type"))
}

// Ensure that value is of type T and return bool
func EnsureTypeBool[T any](value any) bool {
	_, ok := value.(T)
	return ok
}

// Cast value to type T or panic
func Cast[T any](value any) T {
	return EnsureType[T](value)
}

// Cast all provided values to type T
// Returns result slice of type T and bool if it was successful for if condition
func CastAll[T any](v ...any) (ret []T, casted bool) {
	ret = make([]T, 0)
	for _, vv := range v {
		if val, ok := vv.(T); ok {
			ret = append(ret, val)
			casted = true
		} else {
			ret = make([]T, 0)
			return ret, false
		}
	}
	return
}

// Cast function to AnyFunc
func ToAnyFunc[AT any, RT any](f func(AT) RT) AnyFunc {
    return func(v any) any {
        return f(EnsureType[AT](v))
    }
}

// Cast function to AnyAction
func ToAnyAction[AT any](f func(AT)) AnyAction {
    return func(v any) {
        f(EnsureType[AT](v))
    }
}

// Cast function to AnyFuncResult
func ToAnyFuncResult[AT any, RT any](f func(AT) Result[RT]) AnyFuncResult {
    return func(v any) AnyResult {
		ret := f(EnsureType[AT](v))
		if ret.IsError() {
			return Err[any](ret.UnwrapErr())
		}
		return Ok[any](ret.Unwrap())
    }
}

// Cast function to AnyFuncNullable
func ToAnyFuncNullable[AT any, RT any](f func(AT) Nullable[RT]) AnyFuncNullable {
    return func(v any) AnyNullable {
		ret := f(EnsureType[AT](v))
		if ret.IsError() {
			return Null[any]()
		}
		return NullableValue[any](ret.Unwrap())
    }
}
