package math_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestBitwise(t *testing.T) {
	if BitwiseAnd(2, 3) != 2 {
		t.Error("BitwiseAnd(2, 3) != 2")
	}
	if BitwiseOr(2, 3) != 3 {
		t.Error("BitwiseOr(2, 3) != 3")
	}
	if BitwiseXor(2, 3) != 1 {
		t.Error("BitwiseXor(2, 3) != 1")
	}
	if BitwiseNot(2) != -3 {
		t.Error("BitwiseNot(2) != -3")
	}
	if BitwiseLeftShift(2, 3) != 16 {
		t.Error("BitwiseLeftShift(2, 3) != 16")
	}
	if BitwiseRightShift(32, 3) != 4 {
		t.Error("BitwiseRightShift(32, 3) != 4")
	}
}
