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
