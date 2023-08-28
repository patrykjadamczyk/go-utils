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
