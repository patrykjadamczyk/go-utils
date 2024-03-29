package base_test

import (
	"errors"
	. "github.com/patrykjadamczyk/go-utils/base"
	err "github.com/patrykjadamczyk/go-utils/errors"
	"testing"
)

func TestIsError(t *testing.T) {
	if IsError(errors.New("Hello")) != true {
		t.Error("IsError should be true on real error")
	}
	if IsError(nil) != false {
		t.Error("IsError should be false on nil value")
	}
	if IsError(err.NilError{}) != false {
		t.Error("IsError should be false on NilError Object")
	}
}

func TestPanicIfErrorCase1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("PanicIfError should panic on real error")
		}
	}()
	PanicIfError(errors.New("Hello"))
	t.Error("PanicIfError should panic on real error")
}

func TestPanicIfErrorCase2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfError shouldn't panic on nil value")
		}
	}()
	PanicIfError(nil)
}

func TestPanicIfErrorCase3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfError shouldn't panic on NilError Object")
		}
	}()
	PanicIfError(err.NilError{})
}

func errorableFunction1() error {
	return errors.New("Hello")
}

func errorableFunction2() error {
	return nil
}

func errorableFunction3() error {
	return err.NilError{}
}

func errorableFunction4() (int, error) {
	return 0, errors.New("Hello")
}

func errorableFunction5() (int, error) {
	return 21, nil
}

func errorableFunction6() (int, error) {
	return 12, err.NilError{}
}

func TestPanicIfErrorFromOutputCase1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("PanicIfErrorFromOutput should panic on real error")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction1())
	t.Error("PanicIfErrorFromOutput should panic on real error")
}

func TestPanicIfErrorFromOutputCase2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("PanicIfErrorFromOutput should panic on real error")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction4())
	t.Error("PanicIfErrorFromOutput should panic on real error")
}

func TestPanicIfErrorFromOutputCase3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfErrorFromOutput shouldn't panic on nil value")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction2())
}

func TestPanicIfErrorFromOutputCase4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfErrorFromOutput shouldn't panic on nil value")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction5())
}

func TestPanicIfErrorFromOutputCase5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfErrorFromOutput shouldn't panic on NilError Object")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction3())
}

func TestPanicIfErrorFromOutputCase6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("PanicIfErrorFromOutput shouldn't panic on NilError Object")
		}
	}()
	PanicIfErrorFromOutput(errorableFunction6())
}

func TestAssertionCase1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Assertion should panic on false value")
		}
	}()
	Assert(false)
	t.Error("Assertion should panic on false value")
}

func TestAssertionCase2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Assertion shouldn't panic on true value")
		}
	}()
	Assert(true)
}

func TestAssertionCase3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Assertion shouldn't panic on true value")
		}
	}()
	AssertCustomError(true, errors.New("Hello"))
}

func TestAssertionCase4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Assertion should panic on false value")
		}
	}()
	AssertCustomError(false, errors.New("Hello"))
	t.Error("Assertion should panic on false value")
}

func TestChangePanicIntoError(t *testing.T) {
	t1 := PanicToError(func() int {
		panic(NewError("Hello World"))
	})
	t2 := PanicToError(func() int {
		return 1
	})
	if !t1.IsError() {
		t.Error("PanicToError should return error")
	}
	if t2.IsError() {
		t.Error("PanicToError should not return error")
	}
	if t2.UnwrapOr(2) != 1 {
		t.Error("PanicToError should return 1")
	}
}
