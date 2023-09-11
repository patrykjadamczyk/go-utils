package errors

import (
	"github.com/gookit/goutil/errorx"
)

// Make new error with specified message
func NewError(msg string) error {
	return errorx.New(msg)
}

// Make new error with specified message format and specified arguments for this format
func NewErrorf(format string, args ...interface{}) error {
	return errorx.Newf(format, args...)
}

// Make new error with specified code and message
func NewErrorWithCode(code int, msg string) error {
	return errorx.NewR(code, msg)
}

// Add StackTrace to specified error
func AddStackTrace(err error) error {
	return errorx.WithStack(err)
}
