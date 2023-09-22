package base

// Make Result with ok value
func Ok[T any](value T) Result[T] {
	return MakeOkResult[T](value)
}

// Make Result with error value
func Err[T any](err error) Result[T] {
	return MakeErrorResult[T](err)
}

// Make Result with ok and error values
func ErrWithValue[T any](value T, err error) Result[T] {
	return MakeErrorResultWithValue[T](value, err)
}

// Check if specified value is error
// Supports error, ErrorableGenericResultInterface, bool, int, uint
// If it gets other value it returns false
func CheckErr(possibleError any) bool {
	if pe, ok := possibleError.(error); ok {
		return IsError(pe)
	}
	if pe, ok := possibleError.(ErrorableGenericResultInterface); ok {
		return pe.IsError()
	}
	if pe, ok := possibleError.(bool); ok {
		return pe
	}
	if pe, ok := possibleError.(int); ok {
		return pe != 0
	}
	if pe, ok := possibleError.(uint); ok {
		return pe != 0
	}
	return false
}

// Return a pointer copy of value
func ToPtr[T any](value T) *T {
	return &value
}
