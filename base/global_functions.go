package base

import (
	"reflect"
	"time"

	"github.com/patrykjadamczyk/go-utils/utils"
)

// Make Result with ok value
func Ok[T any](value T) Result[T] {
	return MakeOkResult[T](value)
}

// Make Result with error value
func Err[T any](err error) Result[T] {
	return MakeErrorResult[T](err)
}

// Make Result with ok and error values
func ErrWithValue[T any](value T, err error) Result[T] {
	return MakeErrorResultWithValue[T](value, err)
}

// Make Nullable type with correct value
func Some[T any](value T) Nullable[T] {
	return NullableValue[T](value)
}

// Make Nullable type with null value
func None[T any]() Nullable[T] {
	return Null[T]()
}

// Check if specified value is error
// Supports error, ErrorableGenericResultInterface, bool, int, uint
// If it gets other value it returns false
func CheckErr(possibleError any) bool {
	if pe, ok := possibleError.(error); ok {
		return IsError(pe)
	}
	if pe, ok := possibleError.(ErrorableGenericResultInterface); ok {
		return pe.IsError()
	}
	if pe, ok := possibleError.(bool); ok {
		return pe
	}
	if pe, ok := possibleError.(int); ok {
		return pe != 0
	}
	if pe, ok := possibleError.(uint); ok {
		return pe != 0
	}
	return false
}

// Return a pointer copy of value
func ToPtr[T any](value T) *T {
	return &value
}

// IsNil checks if a value is nil or if it's a reference type with a nil underlying value.
func IsNil(x any) bool {
	defer func() { recover() }() // nolint:errcheck
	return x == nil || reflect.ValueOf(x).IsNil()
}

// Attempt invokes a function N times until it returns valid output. Returning either the caught error or nil. When first argument is less than `1`, the function runs until a successful response is returned.
func Attempt[T any](retries int, action func() Result[T]) (int, Result[T]) {
	var result Result[T]

	for i := 0; retries < 1 || i < retries; i++ {
		result = action()
		if !result.IsError() {
			return i + 1, result
		}
	}

	return retries, result
}

// AttemptWithDelay invokes a function N times until it returns valid output,
// with a pause between each call. Returning either the caught error or nil.
// When first argument is less than `1`, the function runs until a successful
// response is returned.
func AttemptWithDelay[T any](retries int, delay time.Duration, action func() Result[T]) (int, Result[T]) {
	var result Result[T]

	for i := 0; retries < 1 || i < retries; i++ {
		result = action()
		if !result.IsError() {
			return i + 1, result
		}

		if retries < 1 || i+1 < retries {
			time.Sleep(delay)
		}
	}

	return retries, result
}

// Try calls the function and return false in case of error.
func Try(callback any) (ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	result := utils.CallFuncWithoutArgs(callback)
	resultAny := utils.ReflectValueToValue(result...)
	resultVal := ConvertToResultMultiple(resultAny...)

	if resultVal.IsError() {
		ok = false
	}

	return
}

// Identity is a basic Function that returns the original value passed to it, unmodified.
func Identity[T any](v T) T {
	return v
}

// ToAny is a basic Function that returns the original value passed to it, cast to an 'any' type.
func ToAny[T any](v T) any {
	return v
}
