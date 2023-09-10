package base

import (
	"errors"
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestResult(t *testing.T) {
	t1 := MakeResult[int]()
	if t1.IsError() {
		t.Error("Result 1 should not be error")
	}
	if t1.UnwrapOr(2) == 2 {
		t.Error("Result 1 should not be 2")
	}
	t2 := MakeOkResult[int](1)
	if t2.UnwrapOr(2) != 1 {
		t.Error("Result 2 should be 1")
	}
	if t2.IsError() {
		t.Error("Result 2 should not be error")
	}
	t3 := MakeErrorResult[int](errors.New("error"))
	if t3.IsError() != true {
		t.Error("Result 3 should be error")
	}
	t4 := MakeErrorResultWithValue[int](1, errors.New("error"))
	if t4.IsError() != true {
		t.Error("Result 4 should be error")
	}
	if t4.DataValue != 1 {
		t.Error("Result 4 should be 1")
	}
	t5 := MakeResult[int]()
	t5.Ok(1)
	if t5.IsError() {
		t.Error("Result 5 should not be error")
	}
	if t5.UnwrapOr(2) != 1 {
		t.Error("Result 5 should be 1")
	}
	t6 := MakeResult[int]()
	t6.Error(errors.New("error"))
	if t6.IsError() != true {
		t.Error("Result 6 should be error")
	}
}

func TestResultAsUnwrappableInterface(t *testing.T) {
	var n UnwrappableInterface[int] = MakeOkResult[int](1)
	if n.UnwrapOr(2) != 1 {
		t.Error("Result should be 1")
	}
}

func TestResultAsExtendedUnwrappableInterface(t *testing.T) {
	t.Skip("Golang is weird with function types and weird with interfaces like that")
	// var r ExtendedUnwrappableInterface[int] = MakeOkResult[int](1)
	// if r.UnwrapOr(2) != 1 {
	// 	t.Error("Result should implement ExtendedUnwrappableInterface")
	// }
}

func TestResultAndThen(t *testing.T) {
	var r Result[int] = MakeOkResult(1)
	f := func(v int) any {
		return any(v + 1)
	}
	rr := r.AndThen(f)
	if rr != 2 {
		t.Error("Result should be 2")
	}
}

func TestResultAsUnwrappableInterfaceUnwrapCase1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Result should not panic on ok value")
		}
	}()
	var r Result[int] = MakeOkResult[int](1)
	if r.Unwrap() != 1 {
		t.Error("Result should be 1")
	}
}

func TestResultAsUnwrappableInterfaceUnwrapCase2(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Result should panic on error")
		}
	}()
	var r Result[int] = MakeErrorResult[int](errors.New("test2"))
	r.Unwrap()
	t.Error("Result should panic on error")
}

func TestResultAsUnwrappableInterfaceUnwrapOr(t *testing.T) {
	var r Result[int] = MakeOkResult(1)
	if r.UnwrapOr(2) != 1 {
		t.Error("Result should be 1")
	}
	r = MakeErrorResult[int](errors.New("test1"))
	if r.UnwrapOr(2) != 2 {
		t.Error("Result should be 2")
	}
}

func TestResultAsUnwrappableInterfaceUnwrapOrErr(t *testing.T) {
	var r Result[int] = MakeOkResult(1)
	if r.UnwrapOrErr(func(n Result[int]) {
		t.Error("Result should be 1 and not call err function")
	}) != 1 {
		t.Error("Result should be 1")
	}
	r = MakeErrorResult[int](errors.New("test1"))
	if r.UnwrapOrErr(func(n Result[int]) {
	}) != 0 {
		t.Error("Result should be 0")
	}
}

func TestResultAsUnwrappableInterfaceUnwrapAsResultOrErr(t *testing.T) {
	var r Result[int] = MakeOkResult(1)
	if r.UnwrapAsResultOrErr(func(n Result[int]) {
		t.Error("Result should be 1 and not call err function")
	}).UnwrapOr(2) != 1 {
		t.Error("Result should be 1")
	}
	r = MakeErrorResult[int](errors.New("test1"))
	rr := r.UnwrapAsResultOrErr(func(n Result[int]) {
	})
	if rr.IsError() != true {
		t.Error("Result should be error")
	}
	if rr.DataValue != 0 {
		t.Error("Result should be 0")
	}
}

func TestResultAsUnwrappableInterfaceUnwrapWithErr(t *testing.T) {
	var r Result[int] = MakeOkResult(1)
	rr1, err1 := r.UnwrapWithErr()
	if rr1 != 1 {
		t.Error("Result should be 1")
	}
	if err1 != nil {
		t.Error("Result shouldn't have error on ok value")
	}
	r = MakeErrorResult[int](errors.New("test1"))
	rr2, err2 := r.UnwrapWithErr()
	if rr2 != 0 {
		t.Error("Result should be 0")
	}
	if err2 == nil {
		t.Error("Result should have error on err value")
	}
}

func TestResultAsUnwrappableInterfaceExpectCase1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Result should not panic on ok value")
		}
	}()
	var r Result[int] = MakeOkResult[int](1)
	r.Expect(errors.New("test1"))
}

func TestResultAsUnwrappableInterfaceExpectCase2(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Result should panic on err value")
		}
	}()
	var r Result[int] = MakeErrorResult[int](errors.New("test2"))
	r.Expect(errors.New("test2"))
}

func TestGlobalFunctions(t *testing.T) {
	t1 := Ok[int](1)
	if t1.IsError() {
		t.Error("Result 1 should not be error")
	}
	if t1.UnwrapOr(2) != 1 {
		t.Error("Result 1 should be 1")
	}
	t2 := Err[int](errors.New("error"))
	if t2.IsError() != true {
		t.Error("Result 2 should be error")
	}
	t3 := ErrWithValue[int](1, errors.New("error"))
	if t3.IsError() != true {
		t.Error("Result 3 should be error")
	}
	if t3.DataValue != 1 {
		t.Error("Result 3 should be 1")
	}
}

func TestResultUtils(t *testing.T) {
	tf := func(v int) (int, error) {
		if v == 1 {
			return 1, nil
		}
		return 0, errors.New("error")
	}
	t1 := ConvertToResult(1, nil)
	if t1.IsError() {
		t.Error("Result 1 should not be error")
	}
	if t1.UnwrapOr(2) != 1 {
		t.Error("Result 1 should be 1")
	}
	t2 := ConvertToResultMultiple(tf(1))
	t3 := ConvertToResultMultiple(tf(2))
	if t2.IsError() {
		t.Error("Result 2 should not be error")
	}
	t2d := EnsureType[[]any](t2.Unwrap())
	if len(t2d) != 2 {
		t.Error("Result 2 should have 2 element")
	}
	if t2d[0] != 1 {
		t.Error("Result 2 should be 1")
	}
	if t3.IsError() != true {
		t.Error("Result 3 should be error")
	}
	if t3.UnwrapOr(2) != 2 {
		t.Error("Result 3 should be 2")
	}
}
