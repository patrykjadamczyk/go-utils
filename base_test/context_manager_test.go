package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

type TestStruct1 struct {
	state int
}

func (t TestStruct1) Dev() Result[string] {
	return CheckAllResults[string](
		EnsureObjectImplementsInterface[TestStruct1, EnsuredObject](t),
		EnsureObjectImplementsInterface[*TestStruct1, GoContextManagerInterface](&t),
	)
}

func (t *TestStruct1) Close() {
	t.state = 1
}

type TestStruct2 struct {
	state int
}

func (t TestStruct2) Dev() Result[string] {
	return CheckAllResults[string](
		EnsureObjectImplementsInterface[TestStruct2, EnsuredObject](t),
		EnsureObjectImplementsInterface[*TestStruct2, ContextManagerInterface](&t),
	)
}

func (t *TestStruct2) Enter() {
	t.state = 12
}

func (t *TestStruct2) Exit() {
	t.state = 21
}

func TestContextManager(t *testing.T) {
	t1 := TestStruct1{}
	t2 := TestStruct2{}

	if t1.Dev().IsError() {
		t.Error("Result 1 should not be error")
	}
	if t1.Dev().UnwrapOr("error") != "OK" {
		t.Error("Result 1 should be OK")
	}
	if t2.Dev().IsError() {
		t.Error("Result 2 should not be error")
	}
	if t2.Dev().UnwrapOr("error") != "OK" {
		t.Error("Result 2 should be OK")
	}

	WithGoContext(&t1, func (ctx GoContextManagerInterface) int {
		if t1.state != 0 {
			t.Error("t1 should be 0", t1.state)
		}
		return 0
	})
	if t1.state != 1 {
		t.Error("t1 should be 1", t1.state)
	}
	WithContext(&t2, func (ctx ContextManagerInterface) int {
		if t2.state != 12 {
			t.Error("t2 should be 12", t2.state)
		}
		return 0
	})
	if t2.state != 21 {
		t.Error("t2 should be 21", t2.state)
	}
}
