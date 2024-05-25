package math_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestClamp(t *testing.T) {
	if Clamp(21, 1, 10) != 10 {
		t.Error("Clamp(21, 1, 10) != 10")
	}
	if Clamp(-21, 1, 10) != 1 {
		t.Error("Clamp(-21, 1, 10) != 1")
	}
	if Clamp(5, 1, 10) != 5 {
		t.Error("Clamp(5, 1, 10) != 5")
	}
}
