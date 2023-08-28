package base

func Ok[T any](value T) Result[T] {
	return MakeOkResult[T](value)
}

func Err[T any](err error) Result[T] {
	return MakeErrorResult[T](err)
}
