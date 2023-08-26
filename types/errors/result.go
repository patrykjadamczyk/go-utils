package errors

import (
	"github.com/patrykjadamczyk/go-utils/errors"
	"github.com/patrykjadamczyk/go-utils/utils"
)

type Result[T any] struct {
	DataValue  T
	ErrorValue error
}

func (r *Result[T]) Error(err error) Result[T] {
	r.ErrorValue = err
	return *r
}

func (r *Result[T]) Ok(value T) Result[T] {
	r.DataValue = value
	return *r
}

func (r *Result[T]) IsError() bool {
	return utils.IsError(r.ErrorValue)
}

func (r *Result[T]) Expect(err any) {
	if r.IsError() {
		panic(err)
	}
}

func (r *Result[T]) Unwrap() T {
	if r.IsError() {
		panic(errors.UnwrapError{})
	}
	return r.DataValue
}

func (r *Result[T]) UnwrapOr(defaultVal T) T {
	if r.IsError() {
		return defaultVal
	}
	return r.DataValue
}

func (r *Result[T]) AndThen(fn func(T) any) any {
	if r.IsError() {
		return nil
	}
	return fn(r.DataValue)
}
