package base_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestSwitch(t *testing.T) {
	tf := func(v int) int {
		return Switch[int, int](v).Case(func (v int) (int, bool) {
			if v == 1 {
				return 2, true
			}
			return 0, false
		}).Default(func (v int) int {
			return 3
		}).Value
	}
	if tf(1) != 2 {
		t.Error("1 should be 2")
	}
	if tf(2) != 3 {
		t.Error("2 should be 3")
	}
	if tf(3) != 3 {
		t.Error("3 should be 3")
	}
}

func TestTernary(t *testing.T) {
	tf := func(v bool) int {
		return Ternary[int](v, 2, 3)
	}
	if tf(true) != 2 {
		t.Error("true should be 2")
	}
	if tf(false) != 3 {
		t.Error("false should be 3")
	}
	if TernaryNull[error](nil, NewError("test")).Error() != "test" {
		t.Error("nil should be test")
	}
	if TernaryNull[int](1, 2) != 1 {
		t.Error("1 should be 1")
	}
	if TernaryNullable[int](Null[int](), 2) != 2 {
		t.Error("null should be 2")
	}
	if TernaryNullable[int](NullableValue(2), 2) != 2 {
		t.Error("2 should be 2")
	}
	if TernaryResult(Ok(1), 1, 2) != 1 {
		t.Error("1 should be 1")
	}
	if TernaryResult(Err[int](NewError("test")), 1, 2) != 2 {
		t.Error("2 should be 2")
	}
}
