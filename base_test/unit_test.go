package base_test

import (
	"testing"
	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestUnitType(t *testing.T) {
	test1 := Unit[int](1, "seconds")
	if test1.GetValue() != 1 {
		t.Errorf("Expected 1, but got %d", test1.GetValue())
	}
	if test1.GetUnitName() != "seconds" {
		t.Errorf("Expected seconds, but got %s", test1.GetUnitName())
	}
}
