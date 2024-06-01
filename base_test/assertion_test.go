package base_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/tests"
)

type TestAssertionsT1Object struct {
}

func (t1 TestAssertionsT1Object) Assert() {
	Assert(false)
}

func (t1 TestAssertionsT1Object) AssertBool() bool {
	return false
}

func (t1 TestAssertionsT1Object) GuardInit() {
	Assert(false)
}

type TestAssertionsT2Object struct {
}

func (t2 TestAssertionsT2Object) Assert() {
	Assert(true)
}

func (t2 TestAssertionsT2Object) AssertBool() bool {
	return true
}

func (t2 TestAssertionsT2Object) GuardInit() {
	Assert(true)
}

func TestAssertions(t *testing.T) {
	// Test AssertAsErr
	if AssertAsErr(true) != nil {
		t.Error("AssertAsErr(true) should not fail")
	}
	if AssertAsErr(false) == nil {
		t.Error("AssertAsErr(false) should not fail")
	}
	ae := NewError("test")
	if AssertCustomErrorAsErr(true, ae) != nil {
		t.Error("AssertAsErr(true) should not fail")
	}
	if (AssertCustomErrorAsErr(false, ae) == nil ||
		AssertCustomErrorAsErr(false, ae).Error() != ae.Error()) {
		t.Error("AssertAsErr(false) should not fail")
	}
	// Test Assert
	tests.TestIfNotPanics(t, func() {
		Assert(true)
	}, "Assert on true should not fail")
	tests.TestIfPanics(t, func() {
		Assert(false)
	}, "Assert on false should fail")

	// Test AssertThat.IsTrue
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsTrue(true)
	}, "AssertThat.IsTrue should not fail on true")
	tests.TestIfPanics(t, func() {
		AssertThat().IsTrue(false)
	}, "AssertThat.IsTrue should fail on false")

	// Test AssertThat.IsFalse
	tests.TestIfPanics(t, func() {
		AssertThat().IsFalse(true)
	}, "AssertThat.IsFalse should fail on true")
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsFalse(false)
	}, "AssertThat.IsFalse should not fail on false")

	// Test AssertThat.ErrorIsNil
	tests.TestIfNotPanics(t, func() {
		AssertThat().ErrorIsNil(nil)
	}, "AssertThat.ErrorIsNil should not fail on nil")
	tests.TestIfPanics(t, func() {
		AssertThat().ErrorIsNil(NewError("Hello"))
	}, "AssertThat.ErrorIsNil should fail on not nil")

	// Test AssertThat.ErrorIsNotNil
	tests.TestIfPanics(t, func() {
		AssertThat().ErrorIsNotNil(nil)
	}, "AssertThat.ErrorIsNotNil should fail on nil")
	tests.TestIfNotPanics(t, func() {
		AssertThat().ErrorIsNotNil(NewError("Hello"))
	}, "AssertThat.ErrorIsNotNil should not fail on not nil")

	// Test AssertThat.IsNil
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsNil(nil)
	}, "AssertThat.IsNil should not fail on nil")
	tests.TestIfPanics(t, func() {
		AssertThat().IsNil(0)
	}, "AssertThat.IsNil should fail on not nil")

	// Test AssertThat.IsNotNil
	tests.TestIfPanics(t, func() {
		AssertThat().IsNotNil(nil)
	}, "AssertThat.IsNotNil should fail on nil")
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsNotNil("Hello")
	}, "AssertThat.IsNotNil should not fail on not nil")

	// Test AssertThat.IsNull
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsNull(nil)
	}, "AssertThat.IsNull should not fail on nil")
	tests.TestIfPanics(t, func() {
		AssertThat().IsNull("Hello")
	}, "AssertThat.IsNull should fail on not nil")

	// Test AssertThat.IsNotNull
	tests.TestIfPanics(t, func() {
		AssertThat().IsNotNull(nil)
	}, "AssertThat.IsNotNull should fail on nil")
	tests.TestIfNotPanics(t, func() {
		AssertThat().IsNotNull("Hello")
	}, "AssertThat.IsNotNull should not fail on not nil")

	// Test AssertThat.Contains
	tests.TestIfNotPanics(t, func() {
		AssertThat().Contains("Hello, World!", "Hello")
	}, "AssertThat.Contains should not fail on substring present")
	tests.TestIfPanics(t, func() {
		AssertThat().Contains("Hello, World!", "Foo")
	}, "AssertThat.Contains should fail on substring not present")

	// Test AssertThat.NotContains
	tests.TestIfNotPanics(t, func() {
		AssertThat().NotContains("Hello, World!", "Foo")
	}, "AssertThat.NotContains should not fail on substring not present")
	tests.TestIfPanics(t, func() {
		AssertThat().NotContains("Hello, World!", "Hello")
	}, "AssertThat.NotContains should fail on substring present")

	// Test AssertThat.AssertableObjectIsValid
	tests.TestIfPanics(t, func() {
		AssertThat().AssertableObjectIsValid(TestAssertionsT1Object{})
	}, "AssertThat.AssertableObjectIsValid should fail on invalid object")
	tests.TestIfNotPanics(t, func() {
		AssertThat().AssertableObjectIsValid(TestAssertionsT2Object{})
	}, "AssertThat.AssertableObjectIsValid should not fail on valid object")

	// Test AssertThat.AssertableBoolObjectIsValid
	tests.TestIfNotPanics(t, func() {
		AssertThat().AssertableBoolObjectIsValid(TestAssertionsT2Object{})
	}, "AssertThat.AssertableBoolObjectIsValid should not fail on valid object")

	tests.TestIfPanics(t, func() {
		AssertThat().AssertableBoolObjectIsValid(TestAssertionsT1Object{})
	}, "AssertThat.AssertableBoolObjectIsValid should fail on invalid object")

	// Test AssertThat.GuardedObjectIsCorrectlyInitialized
	tests.TestIfNotPanics(t, func() {
		AssertThat().GuardedObjectIsCorrectlyInitialized(TestAssertionsT2Object{})
	}, "AssertThat.GuardedObjectIsCorrectlyInitialized should not fail on correct initialization")
	tests.TestIfPanics(t, func() {
		AssertThat().GuardedObjectIsCorrectlyInitialized(TestAssertionsT1Object{})
	}, "AssertThat.GuardedObjectIsCorrectlyInitialized should fail on incorrect initialization")

	// Test AssertThat.Panics
	tests.TestIfNotPanics(t, func() {
		AssertThat().Panics(func() {
			panic("test panic")
		})
	}, "AssertThat.Panics should not fail on panic")
	tests.TestIfPanics(t, func() {
		AssertThat().Panics(func() {})
	}, "AssertThat.Panics should fail on non panic")
}

func TestAssertionsTyped(t *testing.T) {
	// Can call IsNull method with a null value of type T without raising an error
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped[error]().IsNull(nil)
	}, "AssertThatTyped[int].IsNull: should not fail on null value")

	// Can call IsNotNull method with a non-null value of type T without raising an error
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped[int]().IsNotNull(42)
	}, "AssertThatTyped[int].IsNotNull: should not fail on non-null value")

	// Calling IsNull method with a non-null value of type T raises an AssertionError
	tests.TestIfPanics(t, func() {
		AssertThatTyped[int]().IsNull(42)
	}, "AssertThatTyped[int].IsNull: should fail on non-null value")

	// Calling IsNotNull method with a null value of type T raises an AssertionError
	tests.TestIfPanics(t, func() {
		AssertThatTyped[error]().IsNotNull(nil)
	}, "AssertThatTyped[int].IsNotNull: should fail on null value")

	// Calling IsOfTypeT method with a value of a different type than T raises an AssertionError
	tests.TestIfPanics(t, func() {
		AssertThatTyped[int]().IsOfTypeT("test")
	}, "AssertThatTyped[int].IsOfTypeT: should fail on different type")

	// Can call IsOfTypeT method with a value of type T without raising an error
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped[int]().IsOfTypeT(10)
	}, "AssertThatTyped[int].IsOfTypeT: should not fail on type T value")

	// Can call IsOfTypeT method with a value of type T without raising an error
	tests.TestIfNotPanics(t, func() {
		var v int
		AssertThatTyped[int]().IsOfTypeT(v)
	}, "AssertThatTyped[int].IsOfTypeT: should not fail on type T value")

	// Can create an instance of assertThatTyped for a given type T and call all its methods without raising an error
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped[error]().IsNull(nil)
		AssertThatTyped[string]().IsNotNull("test")
		AssertThatTyped[string]().IsOfTypeT("test")
	}, "AssertThatTyped[string]: should not fail on all methods")

	// Calling IsOfTypeT method with a value of type T or a subtype of T does not raise an AssertionError
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped[int]().IsOfTypeT(10)
	}, "AssertThatTyped[int].IsOfTypeT: should not fail on type T or subtype")
}

func TestAssertionsTyped2(t *testing.T) {
	type testInt = int
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped2[testInt, int]().IsInstanceOf(10)
	}, "AssertThatTyped2[testInt,int].IsInstanceOf: should not fail on subtype")
	tests.TestIfPanics(t, func() {
		AssertThatTyped2[testInt, int]().IsNotInstanceOf(10)
	}, "AssertThatTyped2[testInt,int].IsNotInstanceOf should fail on subtype")
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped2[AssertableInterface, TestAssertionsT1Object]().ImplementsT1(TestAssertionsT1Object{})
	}, "AssertThatTyped2[AssertableInterface,T1].ImplementsT1: should not fail on correct interface")
	tests.TestIfPanics(t, func() {
		AssertThatTyped2[AssertableInterface, TestAssertionsT1Object]().NotImplementsT1(TestAssertionsT1Object{})
	}, "AssertThatTyped2[AssertableInterface,T1].NotImplementsT1: should fail on correct interface")
}

func TestAssertionsTypedIsOfOneTypes(t *testing.T) {
	// AssertThatTyped2
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped2[AssertableInterface, int]().IsOfOneTypes(10)
	}, "AssertThatTyped2.IsOfOneTypes: should work on correct types")
	tests.TestIfPanics(t, func() {
		AssertThatTyped2[AssertableInterface, AssertableBoolInterface]().IsOfOneTypes(10)
	}, "AssertThatTyped2.IsOfOneTypes: should fail on incorrect types")
	// AssertThatTyped3
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped3[AssertableInterface, AssertableBoolInterface, int]().IsOfOneTypes(10)
	}, "AssertThatTyped3.IsOfOneTypes: should work on correct types")
	tests.TestIfPanics(t, func() {
		AssertThatTyped3[AssertableInterface, AssertableBoolInterface, UnwrappableInterface[int]]().IsOfOneTypes(10)
	}, "AssertThatTyped3.IsOfOneTypes: should fail on incorrect types")
	// AssertThatTyped4
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped4[AssertableInterface, AssertableBoolInterface, UnwrappableInterface[int], int]().IsOfOneTypes(10)
	}, "AssertThatTyped4.IsOfOneTypes: should work on correct types")
	tests.TestIfPanics(t, func() {
		AssertThatTyped4[AssertableInterface, AssertableBoolInterface, UnwrappableInterface[int], UnwrappableInterface[string]]().IsOfOneTypes(10)
	}, "AssertThatTyped4.IsOfOneTypes: should fail on incorrect types")
	// AssertThatTyped5
	tests.TestIfNotPanics(t, func() {
		AssertThatTyped5[AssertableInterface, AssertableBoolInterface, UnwrappableInterface[int], UnwrappableInterface[string], int]().IsOfOneTypes(10)
	}, "AssertThatTyped5.IsOfOneTypes: should work on correct types")
	tests.TestIfPanics(t, func() {
		AssertThatTyped5[AssertableInterface, AssertableBoolInterface, UnwrappableInterface[int], UnwrappableInterface[string], UnitValue[int]]().IsOfOneTypes(10)
	}, "AssertThatTyped5.IsOfOneTypes: should fail on incorrect types")
}
