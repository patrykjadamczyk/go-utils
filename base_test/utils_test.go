package base_test

import (
	"testing"
	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestRange(t *testing.T) {
	t1 := 0
	for i := range Range(10) {
		t1 += i
	}
	if t1 != 45 {
		t.Error("Range should work", t1)
	}
	t2 := 0
	for _, i2 := range StartEndRange(5, 10) {
		t2 += i2
	}
	if t2 != 45 {
		t.Error("StartEndRange should work", t2)
	}
	t3 := 0
	for _, i3 := range ComplexRange(5, 10, 2) {
		t3 += i3
	}
	if t3 != 21 {
		t.Error("ComplexRange should work", t3)
	}
	t4 := 0
	for _, i4 := range ComplexNonInclusiveRange(5, 10, 2) {
		t4 += i4
	}
	if t4 != 21 {
		t.Error("ComplexNonInclusiveRange should work", t4)
	}
	t5 := 0
	for _, i5 := range StartEndNonInclusiveRange(5, 10) {
		t5 += i5
	}
	if t5 != 35 {
		t.Error("StartEndNonInclusiveRange should work", t5)
	}
}
