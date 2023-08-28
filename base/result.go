package base

import (
	"github.com/patrykjadamczyk/go-utils/errors"
	"github.com/patrykjadamczyk/go-utils/utils"
)

type Result[T any] struct {
	DataValue  T
	ErrorValue error
}

// Make Result
func MakeResult[T any]() Result[T] {
	return Result[T]{}
}

// Make Ok Result
func MakeOkResult[T any](value T) Result[T] {
	r := MakeResult[T]()
	return r.Ok(value)
}

// Make Error Result
func MakeErrorResult[T any](err error) Result[T] {
	r := MakeResult[T]()
	return r.Error(err)
}

// Make Error Result with additional data
func MakeErrorResultWithValue[T any](value T, err error) Result[T] {
	return Result[T]{DataValue: value, ErrorValue: err}
}

// Get Pointer of Result
func (r Result[T]) ToPointer() *Result[T] {
	return &r
}

// Get Result from Pointer
func (r *Result[T]) ToResult() Result[T] {
	return *r
}

// Ok Result
func (r *Result[T]) Ok(value T) Result[T] {
	r.DataValue = value
	return *r
}

// Error Result
func (r *Result[T]) Error(err error) Result[T] {
	r.ErrorValue = err
	return *r
}

// Run function fn when Result is Ok
func (r Result[T]) AndThen(fn func(T) any) any {
	if r.IsError() {
		return nil
	}
	return fn(r.DataValue)
}

// Implementation of ErrorableGenericResultInterface

// Check if Result has error
func (r Result[T]) IsError() bool {
	return utils.IsError(r.ErrorValue)
}

// If result has error get error otherwise return nil
func (r Result[T]) GetError() error {
	return r.ErrorValue
}

// Implementation of ErrorableGenericResultInterface

// Get underlying value or panic if error
func (r Result[T]) Unwrap() T {
	if r.IsError() {
		panic(errors.UnwrapError{})
	}
	return r.DataValue
}

// Get underlying value or return default value if error is found
func (r Result[T]) UnwrapOr(defaultVal T) T {
	if r.IsError() {
		return defaultVal
	}
	return r.DataValue
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (r Result[T]) UnwrapOrErr(f func(Result[T])) T {
	if r.IsError() {
		f(r)
	}
	return r.DataValue
}

// Unwrap value and error separately (Result -> Go normal returns)
func (r Result[T]) UnwrapWithErr() (T, error) {
	return r.DataValue, r.ErrorValue
}

// Expect correct value if error is found panic with specified message
func (r Result[T]) Expect(err any) {
	if r.IsError() {
		panic(err)
	}
}
