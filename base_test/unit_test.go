package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestUnitType(t *testing.T) {
	test1 := Unit[int](1, "seconds")
	if test1.GetValue() != 1 {
		t.Errorf("Expected 1, but got %d", test1.GetValue())
	}
	if test1.GetUnitName() != "seconds" {
		t.Errorf("Expected seconds, but got %s", test1.GetUnitName())
	}
	tf2 := MakeUnit[int]("minutes")
	test2 := tf2(2)
	if test2.GetValue() != 2 {
		t.Errorf("Expected 2, but got %d", test2.GetValue())
	}
	if test2.GetUnitName() != "minutes" {
		t.Errorf("Expected minutes, but got %s", test2.GetUnitName())
	}
	tf3 := func(v any) Result[UnitValue[int]] {
		value := NormalizeUnit[int]("hours")(v)
		if value.IsError() {
			return Err[UnitValue[int]](value.GetError())
		}
		return Ok[UnitValue[int]](value.Unwrap())
	}
	test3 := tf3(3)
	if test3.IsError() {
		t.Errorf("Expected Ok, but got %s", test3.GetError())
	}
	if test3.Unwrap().GetValue() != 3 {
		t.Errorf("Expected 3, but got %d", test3.Unwrap().GetValue())
	}
	if test3.Unwrap().GetUnitName() != "hours" {
		t.Errorf("Expected hours, but got %s", test3.Unwrap().GetUnitName())
	}
}
