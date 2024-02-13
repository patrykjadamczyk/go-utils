package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestVariadic(t *testing.T) {
	ta := []int{1, 2, 3}
	t1 := PackToArray(ConvertToAnyArray(ta)...)
	if t1[0] != 1 || t1[1] != 2 || t1[2] != 3 {
		t.Error("PackToArray should work")
	}
	var t2a int
	var t2b int
	var t2c int
	ExplodeTyped(ta, &t2a, &t2b, &t2c)
	if t2a != 1 || t2b != 2 || t2c != 3 {
		t.Error("ExplodeTyped should work")
	}
	var t3a any
	var t3b any
	var t3c any
	Explode(t1, &t3a, &t3b, &t3c)
	if t3a != 1 || t3b != 2 || t3c != 3 {
		t.Error("Explode should work")
	}
}
