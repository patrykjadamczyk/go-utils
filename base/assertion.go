package base

import (
	"github.com/patrykjadamczyk/go-utils/errors"
)

// Assert that condition is true if it's not panic with AssertionError
func Assert(cond bool) {
	if !cond {
		panic(ExpandError(errors.AssertionError{}))
	}
}

// Assert that condition is true if it's not panic with specified custom error
func AssertCustomError(cond bool, err error) {
	if !cond {
		panic(ExpandError(err))
	}
}

