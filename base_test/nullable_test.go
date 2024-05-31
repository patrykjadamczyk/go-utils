package base_test

import (
	"errors"
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestNullableShouldBeUnwrappable(t *testing.T) {
	var n UnwrappableInterface[int] = NullableValue(1)
	if n.UnwrapOr(2) != 1 {
		t.Error("Nullable should implement UnwrappableInterface")
	}
}

func TestNullableShouldBeExtendedUnwrappable(t *testing.T) {
	t.Skip("Golang is weird with function types and weird with interfaces like that")
	// var n ExtendedUnwrappableInterface[int] = NullableValue(1)
	// if n.UnwrapOr(2) != 1 {
	// t.Error("Nullable should implement ExtendedUnwrappableInterface")
	// }
}

func TestNullableShouldBeErrorable(t *testing.T) {
	var n ErrorableGenericResultInterface = NullableValue(1)
	if n.IsError() {
		t.Error("Nullable should implement ErrorableGenericResultInterface")
	}
}

func TestNullableValue(t *testing.T) {
	var n = NullableValue(1)
	if n.ValueOrZero() != 1 {
		t.Error("Nullable should be 1")
	}
	n = Null[int]()
	if n.IsZero() != true {
		t.Error("Nullable should be NULL")
	}
	if n.ValueOrZero() != 0 {
		t.Error("Nullable should be 0")
	}
	n = NullableValue(0)
	if n.IsZero() != true {
		t.Error("Nullable should be NULL")
	}
	if n.ValueOrZero() != 0 {
		t.Error("Nullable should be 0")
	}
	n1 := NullableValue(1)
	n2 := NullableValue(2)
	n3 := NullableValue(1)

	if n1.Equal(n2) != false {
		t.Error("Nullable should not be equal")
	}
	if n1.Equal(n3) != true {
		t.Error("Nullable should be equal")
	}

	if n1.GoString() != "Nullable[int]{Data:1,Valid:true}" {
		t.Error("Nullable should have correct GoString")
	}
}

func TestNullableAndThen(t *testing.T) {
	var n = NullableValue(1)
	f := func(v int) any {
		return any(v + 1)
	}
	nr := n.AndThen(f)
	if nr != 2 {
		t.Error("Nullable should be 2")
	}
}

func TestNullableAsUnwrappableInterfaceUnwrapCase1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Nullable should not panic on correct value")
		}
	}()
	var n = NullableValue(1)
	if n.Unwrap() != 1 {
		t.Error("Nullable should be 1")
	}
}

func TestNullableAsUnwrappableInterfaceUnwrapCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Nullable should panic on incorrect value")
		}
	}()
	var n = Null[int]()
	n.Unwrap()
	t.Error("Nullable should panic on nil value unwrapping")
}

func TestNullableAsUnwrappableInterfaceUnwrapErrCase1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Nullable should not panic on incorrect value")
		}
	}()
	var n = Null[int]()
	_ = n.UnwrapErr()
}

func TestNullableAsUnwrappableInterfaceUnwrapErrCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Nullable should panic on correct value")
		}
	}()
	var n = NullableValue(1)
	_ = n.UnwrapErr()
	t.Error("Nullable should panic on correct value unwrapping")
}

func TestNullableAsUnwrappableInterfaceUnwrapOr(t *testing.T) {
	var n = NullableValue(1)
	if n.UnwrapOr(2) != 1 {
		t.Error("Nullable should be 1")
	}
	n = Null[int]()
	if n.UnwrapOr(2) != 2 {
		t.Error("Nullable should be 2")
	}
}

func TestNullableAsUnwrappableInterfaceUnwrapOrErr(t *testing.T) {
	var n = NullableValue(1)
	if n.UnwrapOrErr(func(_ Nullable[int]) {
		t.Error("Nullable should be 1 and not call err function")
	}) != 1 {
		t.Error("Nullable should be 1")
	}
	n = Null[int]()
	if n.UnwrapOrErr(func(n Nullable[int]) {
	}) != 0 {
		t.Error("Nullable should be 0")
	}
}

func TestNullableAsUnwrappableInterfaceUnwrapAsResultOrErr(t *testing.T) {
	var n = NullableValue(1)
	if n.UnwrapAsResultOrErr(func(_ Nullable[int]) {
		t.Error("Nullable should be 1 and not call err function")
	}).UnwrapOr(2) != 1 {
		t.Error("Nullable should be 1")
	}
	n = Null[int]()
	nr := n.UnwrapAsResultOrErr(func(n Nullable[int]) {
	})
	if nr.IsError() != true {
		t.Error("Nullable should be error")
	}
	if nr.DataValue != 0 {
		t.Error("Nullable should be 0")
	}
}

func TestNullableAsUnwrappableInterfaceUnwrapWithErr(t *testing.T) {
	var n = NullableValue(1)
	nr1, err1 := n.UnwrapWithErr()
	if nr1 != 1 {
		t.Error("Nullable should be 1")
	}
	if err1 != nil {
		t.Error("Nullable shouldn't have error on correct value")
	}
	n = Null[int]()
	nr2, err2 := n.UnwrapWithErr()
	if nr2 != 0 {
		t.Error("Nullable should be 0")
	}
	if err2 == nil {
		t.Error("Nullable should have error on incorrect value")
	}
}

func TestNullableAsUnwrappableInterfaceExpectCase1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Nullable should not panic on correct value")
		}
	}()
	var n = NullableValue(1)
	n.Expect(errors.New("test1"))
}

func TestNullableAsUnwrappableInterfaceExpectCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Nullable should panic on incorrect value")
		}
	}()
	var n = Null[int]()
	n.Expect(errors.New("test2"))
}

func TestNullableAsUnwrappableInterfaceExpectErrCase1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Nullable should not panic on correct value")
		}
	}()
	var n = Null[int]()
	n.ExpectErr(errors.New("test1"))
}

func TestNullableAsUnwrappableInterfaceExpectErrCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Nullable should panic on incorrect value")
		}
	}()
	var n = NullableValue(1)
	n.ExpectErr(errors.New("test2"))
}

func TestNullableNilInterfaceValueProblem(t *testing.T) {
	var data *byte
	var in any = data
	if !NullableValue(in).IsZero() {
		t.Error("Nullable doesn't handle nil interface value problem")
	}
}
