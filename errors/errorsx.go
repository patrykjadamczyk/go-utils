package errors

import (
	"github.com/gookit/goutil/errorx"
)

func NewError(msg string) error {
	return errorx.New(msg)
}

func NewErrorf(format string, args ...interface{}) error {
	return errorx.Newf(format, args...)
}

func NewErrorWithCode(code int, msg string) error {
	return errorx.NewR(code, msg)
}

func AddStackTrace(err error) error {
	return errorx.WithStack(err)
}
