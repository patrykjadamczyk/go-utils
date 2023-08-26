package utils

type NilError struct {
}

func (NilError) Error() string {
    return "NilError"
}

var NilErrorErr NilError = NilError{}

func IsError(err error) bool {
    if err == nil {
        return false
    }
    switch err.(type) {
    case NilError:
        return false
    default:
        return true
    }
}
