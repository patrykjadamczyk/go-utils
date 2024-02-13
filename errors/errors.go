package errors

type NilError struct {
}

func (NilError) Error() string {
	return "NilError"
}

var NilErrorErr NilError = NilError{}

type UnwrapError struct {
}

func (UnwrapError) Error() string {
	return "UnwrapError: Tried unwrap null value"
}

type NullableError struct {
}

func (NullableError) Error() string {
	return "NullableError: Nullable contains null"
}

type AssertionError struct {
}

func (AssertionError) Error() string {
	return "AssertionError: Assertion failed"
}

type NotImplementedError struct {
}

func (NotImplementedError) Error() string {
	return "NotImplementedError: Not Implemented!"
}

type UnreachableError struct {
}

func (UnreachableError) Error() string {
	return "UnreachableError: Unreachable!"
}
