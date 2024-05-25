package math_test

import (
	"math"
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestNaN(t *testing.T) {
	if !IsNaN(NaN[float32]()) {
		t.Error("Float32 NaN is not NaN")
	}
	if !IsNaN(NaN[float64]()) {
		t.Error("Float64 NaN is not NaN")
	}
	if !math.IsNaN(NaN[float64]()) {
		t.Error("Float64 NaN is not NaN according to math.IsNaN")
	}
}
