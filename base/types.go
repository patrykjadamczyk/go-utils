package base

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
