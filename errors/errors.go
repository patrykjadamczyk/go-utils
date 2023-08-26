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
