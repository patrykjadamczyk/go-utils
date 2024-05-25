package math_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestVector(t *testing.T) {
	vec2 := Vector2[int]{X: 1, Y: 1}
	if vec2.X != 1 || vec2.Y != 1 {
		t.Error("vec2 should be Vector2(1, 1) and is", vec2.ToString())
	}
	if !vec2.Equal(Vector2[int]{X: 1, Y: 1}) {
		t.Error("vec2 should be Vector2(1, 1) and is", vec2.ToString())
	}
	vec2.Add(Vector2[int]{X: 1, Y: 1})
	if !vec2.Equal(Vector2[int]{X: 2, Y: 2}){
        t.Error("vec2 should be Vector2(1, 1) + Vector2(1, 1) = Vector2(2, 2) and is", vec2.ToString())
    }
	vec2.Substract(Vector2[int]{X: 1, Y: 1})
	if !vec2.Equal(Vector2[int]{X: 1, Y: 1}) {
        t.Error("vec2 should be Vector2(2, 2) - Vector2(1, 1) = Vector2(1, 1) and is", vec2.ToString())
    }
	vec2.Multiply(Vector2[int]{X: 2, Y: 2})
	if !vec2.Equal(Vector2[int]{X: 2, Y: 2}) {
        t.Error("vec2 should be Vector2(1, 1) * Vector2(2, 2) = Vector2(2, 2) and is", vec2.ToString())
    }
	vec2.Divide(Vector2[int]{X: 2, Y: 2})
	if !vec2.Equal(Vector2[int]{X: 1, Y: 1}) {
        t.Error("vec2 should be Vector2(2, 2) / Vector2(2, 2) = Vector2(1, 1) and is", vec2.ToString())
    }
	vec2.AddWithinBounds(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: 3, Y: 3})
	if !vec2.Equal(Vector2[int]{X: 0, Y: 0}) {
		t.Error(
			"vec2 should be Vector2(1, 1) + Vector2(3, 3) => in bounds of Vector2(2, 2) = Vector2(0, 0) and is",
			vec2.ToString(),
		)
	}
	vec2.AddWithinBounds(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: -3, Y: -3})
	if !vec2.Equal(Vector2[int]{X: -1, Y: -1}) {
		t.Error(
			"vec2 should be Vector2(1, 1) + Vector2(-3, -3) => in bounds of Vector2(2, 2) = Vector2(-1, -1) and is",
			vec2.ToString(),
		)
	}
	vec2.AddWithinBoundsAndForceAbsolute(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: -2, Y: -2})
	if !vec2.Equal(Vector2[int]{X: 1, Y: 1}) {
		t.Error(
			"vec2 should be Vector2(-1, -1) + Vector2(-2, -2) => in bounds of Vector2(2, 2) => Abs() = Vector2(1, 1) and is",
			vec2.ToString(),
		)
	}
}
