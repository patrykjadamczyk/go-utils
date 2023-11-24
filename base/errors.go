package base

import (
	err "errors"
	"github.com/patrykjadamczyk/go-utils/errors"
)

// Check if specified value is error or it's nil or NilError
func IsError(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case errors.NilError:
		return false
	default:
		return true
	}
}

// Panic if value provided is error
func PanicIfError(err error) {
	if IsError(err) {
		panic(err)
	}
}

// Panic if Error from specified function output has error
func PanicIfErrorFromOutput(returns ...any) {
	for _, v := range returns {
		if err, ok := v.(error); ok {
			PanicIfError(err)
		}
	}
}

// Make New Error
func NewError(msg string) error {
	return errors.NewError(msg)
}

// Expand provided error adding stacktrace
func ExpandError(err error) error {
	return errors.AddStackTrace(err)
}

// Note that something is not yet implement
// This function panics with NotImplementedError
func NotImplemented() {
	panic(ExpandError(errors.NotImplementedError{}))
}

// Note that something is not reachable
// This function panics with UnreachableError
func Unreachable() {
	panic(ExpandError(errors.UnreachableError{}))
}

// Note that something is not supported
// This function panics with NotSupportedError
func NotSupported() {
	panic(ExpandError(err.ErrUnsupported))
}

// Check if 2 errors are equal
func IsErrorEqual(err1 error, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 == nil || err2 == nil {
		return false
	}
	return err.Is(err1, err2)
}

// Change block that could panic into Result with error as value
// Note this function only catches panic errors and strings
func PanicToError[T any](f func() T) Result[T] {
	r := MakeResult[T]()
	pf := func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				if err, ok := recovered.(error); ok {
					r = Err[T](err)
					return
				}
				if errs, ok := recovered.(string); ok {
					r = Err[T](NewError(errs))
					return
				}
				panic(r)
			}
		}()
		fr := f()
		r = Ok[T](fr)
	}
	pf()
	return r
}
