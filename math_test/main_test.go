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

func TestAbs(t *testing.T) {
    if Abs(1) != 1 {
        t.Error("int Abs(1) != 1")
    }
	if Abs[int8](1) != 1 {
        t.Error("int8 Abs(1) != 1")
    }
    if Abs[int16](1) != 1 {
        t.Error("int16 Abs(1) != 1")
    }
    if Abs[int32](1) != 1 {
        t.Error("int32 Abs(1) != 1")
    }
    if Abs[int64](1) != 1 {
        t.Error("int64 Abs(1) != 1")
    }
    if Abs[float32](1) != 1 {
        t.Error("float32 Abs(1) != 1")
    }
    if Abs[float64](1) != 1 {
        t.Error("float64 Abs(1) != 1")
    }
    if Abs(-1) != 1 {
        t.Error("int Abs(-1) != 1")
    }
	if Abs[int8](-1) != 1 {
        t.Error("int8 Abs(-1) != 1")
    }
    if Abs[int16](-1) != 1 {
        t.Error("int16 Abs(-1) != 1")
    }
    if Abs[int32](-1) != 1 {
        t.Error("int32 Abs(-1) != 1")
    }
    if Abs[int64](-1) != 1 {
        t.Error("int64 Abs(-1) != 1")
    }
    if Abs[float32](-1) != 1 {
        t.Error("float32 Abs(-1) != 1")
    }
    if Abs[float64](-1) != 1 {
        t.Error("float64 Abs(-1) != 1")
    }
}

func TestMinMax(t *testing.T) {
	if Min(12, 245, 90) != 12 {
		t.Error("int Min(12, 245, 90) != 12")
	}
	if Min(12, 245, 90, -21) != -21 {
		t.Error("int Min(12, 245, 90, -21) != -21")
	}
	if Max(12, 245, 90) != 245 {
		t.Error("int Max(12, 245, 90) != 245")
	}
	if Max(12, 245, 90, -21) != 245 {
		t.Error("int Max(12, 245, 90, -21) != 245")
	}
}
