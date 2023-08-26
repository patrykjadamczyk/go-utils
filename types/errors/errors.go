package errors

import (
	"github.com/patrykjadamczyk/go-utils/errors"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
    if _, ok := err.(errors.NilError); ok {
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
