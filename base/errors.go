package base

import (
	"github.com/patrykjadamczyk/go-utils/errors"
)

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

func PanicIfError(err error) {
	if IsError(err) {
		panic(err)
	}
}

func PanicIfErrorFromOutput(returns ...any) {
	for _, v := range returns {
		if err, ok := v.(error); ok {
			PanicIfError(err)
		}
	}
}

func NewError(msg string) error {
	return errors.NewError(msg)
}

func ExpandError(err error) error {
	return errors.AddStackTrace(err)
}

func Assert(cond bool) {
	if !cond {
		panic(ExpandError(errors.AssertionError{}))
	}
}

func AssertCustomError(cond bool, err error) {
	if !cond {
		panic(ExpandError(err))
	}
}
