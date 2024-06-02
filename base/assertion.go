package base

import (
	"strings"

	"github.com/patrykjadamczyk/go-utils/errors"
)

// Assert that condition is true if it's not panic with AssertionError
func Assert(cond bool) {
	if !cond {
		panic(ExpandError(errors.AssertionError{}))
	}
}

// Assert that condition is true if it's not then return AssertionError
func AssertAsErr(cond bool) (err error) {
	if !cond {
		err = ExpandError(errors.AssertionError{})
	}
	return
}

// Assert that condition is true if it's not panic with specified custom error
func AssertCustomError(cond bool, err error) {
	if !cond {
		panic(ExpandError(err))
	}
}

// Assert that condition is true if it's not panic with specified custom error
func AssertCustomErrorString(cond bool, err string) {
	if !cond {
		panic(ExpandError(NewError(err)))
	}
}

// Assert that condition is true if it's not then return AssertionError
func AssertCustomErrorAsErr(cond bool, err error) (rerr error) {
	if !cond {
		rerr = ExpandError(err)
	}
	return
}

// Assert that condition is true if it's not then return AssertionError
func AssertCustomErrorStringAsErr(cond bool, err string) (rerr error) {
	if !cond {
		rerr = ExpandError(NewError(err))
	}
	return
}

// assertThat is a container for basic assertion methods
// that don't require type parameters if some method are more generic it uses any
type assertThat struct{}

func AssertThat() assertThat {
	return assertThat{}
}

// Assert that condition is true
func (a assertThat) IsTrue(cond bool) {
	Assert(cond)
}

// Assert that condition is false
func (a assertThat) IsFalse(cond bool) {
	Assert(!cond)
}

// Assert that error is nil
func (a assertThat) ErrorIsNil(err error) {
	Assert(err == nil)
}

// Assert that error is not nil
func (a assertThat) ErrorIsNotNil(err error) {
	Assert(err != nil)
}

// Assert that value is nil
func (a assertThat) IsNil(v any) {
	Assert(v == nil)
}

// Assert that value is not nil
func (a assertThat) IsNotNil(v any) {
	Assert(v != nil)
}

// Assert that value is null (using Nullable)
func (a assertThat) IsNull(v any) {
	Assert(NullableValue(v).IsError())
}

// Assert that value is not null (using Nullable)
func (a assertThat) IsNotNull(v any) {
	Assert(!NullableValue(v).IsError())
}

// Assert that string contains specified other string
func (a assertThat) Contains(testValue string, containString string) {
	Assert(strings.Contains(testValue, containString))
}

// Assert that string does not contain specified other string
func (a assertThat) NotContains(testValue string, containString string) {
	Assert(!strings.Contains(testValue, containString))
}

// Assert that assertable object is valid
func (a assertThat) AssertableObjectIsValid(v AssertableInterface) {
	v.Assert()
}

// Assert that assertable bool object is valid
func (a assertThat) AssertableBoolObjectIsValid(v AssertableBoolInterface) {
	Assert(v.AssertBool())
}

// Assert that GuardedObject is correctly initialized
func (a assertThat) GuardedObjectIsCorrectlyInitialized(v GuardedObjectInterface) {
	v.GuardInit()
}

// Assert that function panics
func (a assertThat) panics(f func()) (ret bool) {
	defer func() {
		if r := recover(); r != nil {
			ret = true
		} else {
			ret = false
		}
	}()
	f()
	ret = false
	return
}

// Assert that function panics
func (a assertThat) Panics(f func()) {
	Assert(a.panics(f))
}

// Assert that function does not panic
func (a assertThat) NotPanics(f func()) {
	Assert(!a.panics(f))
}

// assertThatTyped is a container for assertion methods that need one type parameter
type assertThatTyped[T any] struct{}

func AssertThatTyped[T any]() assertThatTyped[T] {
	return assertThatTyped[T]{}
}

// Assert that typed value is null
func (a assertThatTyped[T]) IsNull(v T) {
	Assert(NullableValue(v).IsError())
}

// Assert that typed value is not null
func (a assertThatTyped[T]) IsNotNull(v T) {
	Assert(!NullableValue(v).IsError())
}

// Assert that value is of type T
func (a assertThatTyped[T]) IsOfTypeT(v any) {
	Assert(EnsureTypeBool[T](v))
}

// AssertThatType is a container for assertion methods that need two type parameter
type assertThatTyped2[T1 any, T2 any] struct{}

func AssertThatTyped2[T1 any, T2 any]() assertThatTyped2[T1, T2] {
	return assertThatTyped2[T1, T2]{}
}

// Assert that specified object of type T2 is instance of type T1
func (a assertThatTyped2[T1, T2]) IsInstanceOf(v T2) {
	Assert(EnsureTypeBool[T1](v))
}

// Assert that specified object of type T2 is not instance of type T1
func (a assertThatTyped2[T1, T2]) IsNotInstanceOf(v T2) {
	Assert(!EnsureTypeBool[T1](v))
}

// Assert that specified object of type T2 implements T1
func (a assertThatTyped2[T1, T2]) ImplementsT1(v T2) {
	Assert(!EnsureObjectImplementsInterface[T2, T1](v).IsError())
}

// Assert that specified object of type T2 does not implements T1
func (a assertThatTyped2[T1, T2]) NotImplementsT1(v T2) {
	Assert(EnsureObjectImplementsInterface[T2, T1](v).IsError())
}

// Assert that specified object is of one of types specified in type parameters
func (a assertThatTyped2[T1, T2]) IsOfOneTypes(v any) {
	Assert(EnsureTypeBool[T1](v) || EnsureTypeBool[T2](v))
}

// AssertThatType is a container for assertion methods that need two type parameter
type assertThatTyped3[T1 any, T2 any, T3 any] struct{}

func AssertThatTyped3[T1 any, T2 any, T3 any]() assertThatTyped3[T1, T2, T3] {
	return assertThatTyped3[T1, T2, T3]{}
}

// Assert that specified object is of one of types specified in type parameters
func (a assertThatTyped3[T1, T2, T3]) IsOfOneTypes(v any) {
	Assert(EnsureTypeBool[T1](v) || EnsureTypeBool[T2](v) || EnsureTypeBool[T3](v))
}

// AssertThatType is a container for assertion methods that need two type parameter
type assertThatTyped4[T1 any, T2 any, T3 any, T4 any] struct{}

func AssertThatTyped4[T1 any, T2 any, T3 any, T4 any]() assertThatTyped4[T1, T2, T3, T4] {
	return assertThatTyped4[T1, T2, T3, T4]{}
}

// Assert that specified object is of one of types specified in type parameters
func (a assertThatTyped4[T1, T2, T3, T4]) IsOfOneTypes(v any) {
	Assert(EnsureTypeBool[T1](v) || EnsureTypeBool[T2](v) || EnsureTypeBool[T3](v) || EnsureTypeBool[T4](v))
}

// AssertThatType is a container for assertion methods that need two type parameter
type assertThatTyped5[T1 any, T2 any, T3 any, T4 any, T5 any] struct{}

func AssertThatTyped5[T1 any, T2 any, T3 any, T4 any, T5 any]() assertThatTyped5[T1, T2, T3, T4, T5] {
	return assertThatTyped5[T1, T2, T3, T4, T5]{}
}

// Assert that specified object is of one of types specified in type parameters
func (a assertThatTyped5[T1, T2, T3, T4, T5]) IsOfOneTypes(v any) {
	Assert(EnsureTypeBool[T1](v) || EnsureTypeBool[T2](v) || EnsureTypeBool[T3](v) || EnsureTypeBool[T4](v) || EnsureTypeBool[T5](v))
}
