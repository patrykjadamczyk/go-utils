package base

func Ok[T any](value T) Result[T] {
	return MakeOkResult[T](value)
}

func Err[T any](err error) Result[T] {
	return MakeErrorResult[T](err)
}

func ErrWithValue[T any](value T, err error) Result[T] {
	return MakeErrorResultWithValue[T](value, err)
}

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
