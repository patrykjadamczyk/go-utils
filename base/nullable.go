package base

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/patrykjadamczyk/go-utils/errors"
	"github.com/patrykjadamczyk/go-utils/utils"
)

// Nullable represents data that also can be NULL
// Implements UnwrappableInterface
type Nullable[T any] struct {
	Data  T
	Valid bool
}

// Value Create a Nullable from a value
func NullableValue[T any](value T) Nullable[T] {
	if any(value) == nil {
		return Nullable[T]{Valid: false}
	}
	switch any(value).(type) {
	case errors.NilError:
		return Nullable[T]{Valid: false}
	default:
		return Nullable[T]{Data: value, Valid: true}
	}
}

// ValueFromPointer Create a Nullable from a pointer
func NullableValueFromPointer[T any](value *T) Nullable[T] {
	if value == nil {
		return Nullable[T]{Valid: false}
	}
	return NullableValue(*value)
}

// Null Create a Nullable that is NULL with type
func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

// Null value
func NullValue[T any]() T {
	return Null[T]().ValueOrZero()
}

// ValueOrZero Get Value, or default zero value if it is NULL
func (n Nullable[T]) ValueOrZero() T {
	if !n.Valid {
		var ref T
		return ref
	}
	return n.Data
}

// Check for case of nil value interface
func (n Nullable[T]) isNilString() (result bool) {
	defer func() {
		if r := recover(); r == nil {
			result = true
		}
	}()
	return fmt.Sprintf("%v", n.Data) == utils.FmtSprintfVNil
}

func (n Nullable[T]) IsZero() bool {
	if !n.Valid {
		return true
	}
	var ref T
	return any(ref) == any(n.Data) || n.isNilString()
}

// Equal Check if this Nullable is equal to another Nullable
func (n Nullable[T]) Equal(other Nullable[T]) bool {
	switch any(n.Data).(type) {
	case time.Time:
		nValue := any(n.Data).(time.Time)
		otherValue := any(other.Data).(time.Time)
		return n.Valid == other.Valid && (!n.Valid || nValue.Equal(otherValue))
	}
	return n.ExactEqual(other)
}

// ExactEqual Check if this Nullable is exact equal to another Nullable, never using intern Equal method to check equality
func (n Nullable[T]) ExactEqual(other Nullable[T]) bool {
	return n.Valid == other.Valid && (!n.Valid || any(n.Data) == any(other.Data))
}

// String Convert value to string
func (n Nullable[T]) String() string {
	return fmt.Sprintf("%s", any(n.Data))
}

// Get Go String
func (n Nullable[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"Nullable[%T]{Data:%#v,Valid:%#v}",
		ref,
		n.Data,
		n.Valid,
	)
}

// Get underlying value
func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Data, nil
}

// Run function fn when value is not NULL
func (n Nullable[T]) AndThen(fn func(T) any) any {
	if !n.Valid {
		return nil
	}
	return fn(n.Data)
}

// UnwrappableInterface Implementation

// Get underlying value or panic if error
func (n Nullable[T]) Unwrap() T {
	if !n.Valid {
		panic(errors.UnwrapError{})
	}
	return n.Data
}

// Get underlying error or panic if correct value
func (n Nullable[T]) UnwrapErr() error {
	if n.Valid {
		panic(errors.UnwrapError{})
	}
	return errors.NullableError{}
}

// Get underlying value or return default value if error is found
func (n Nullable[T]) UnwrapOr(defaultVal T) T {
	if !n.Valid {
		return defaultVal
	}
	return n.Data
}

// Unwrap but if error occurs run specified function f instead of panic
// Return result of underlying value (having err if it's found)
func (n Nullable[T]) UnwrapOrErr(f func(Nullable[T])) T {
	if !n.Valid {
		f(n)
	}
	return n.ValueOrZero()
}

// Unwraps as result type for if error checks and run specified function f
// instead of panic
// Return result of underlying value (having err if it's found)
func (n Nullable[T]) UnwrapAsResultOrErr(f func(Nullable[T])) Result[T] {
	if !n.Valid {
		f(n)
		return MakeErrorResult[T](errors.UnwrapError{})
	}
	return MakeOkResult[T](n.ValueOrZero())
}

// Unwrap value and error separately (Result -> Go normal returns)
func (n Nullable[T]) UnwrapWithErr() (T, error) {
	var err error = nil
	if !n.Valid {
		err = errors.NullableError{}
	}
	return n.Data, err
}

// Expect correct value if error is found panic with specified message
func (n Nullable[T]) Expect(err any) {
	if !n.Valid {
		panic(err)
	}
}

// Expect error value if error is not found panic with specified message
func (n Nullable[T]) ExpectErr(err any) {
	if n.Valid {
		panic(err)
	}
}

// Implementation of ErrorableGenericResultInterface

// Check if Object has error
func (n Nullable[T]) IsError() bool {
	return !n.Valid
}

// Get Error if Object has error or nil if not
func (n Nullable[T]) GetError() error {
	if n.Valid {
		return nil
	}
	return errors.NullableError{}
}
