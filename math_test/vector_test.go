package math_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestVector2(t *testing.T) {
	vec2 := Vector2[int]{X: 1, Y: 1}
	if vec2.X != 1 || vec2.Y != 1 {
		t.Error("vec2 should be Vector2(1, 1) and is", vec2.ToString())
	}
	if !vec2.Equal(OneVector2[int]()) {
		t.Error("vec2 should be Vector2(1, 1) and is", vec2.ToString())
	}
	vec2.Add(OneVector2[int]())
	if !vec2.Equal(Vector2[int]{X: 2, Y: 2}){
        t.Error("vec2 should be Vector2(1, 1) + Vector2(1, 1) = Vector2(2, 2) and is", vec2.ToString())
    }
	vec2.Substract(OneVector2[int]())
	if !vec2.Equal(OneVector2[int]()) {
        t.Error("vec2 should be Vector2(2, 2) - Vector2(1, 1) = Vector2(1, 1) and is", vec2.ToString())
    }
	vec2.Multiply(Vector2[int]{X: 2, Y: 2})
	if !vec2.Equal(Vector2[int]{X: 2, Y: 2}) {
        t.Error("vec2 should be Vector2(1, 1) * Vector2(2, 2) = Vector2(2, 2) and is", vec2.ToString())
    }
	vec2.Divide(Vector2[int]{X: 2, Y: 2})
	if !vec2.Equal(OneVector2[int]()) {
        t.Error("vec2 should be Vector2(2, 2) / Vector2(2, 2) = Vector2(1, 1) and is", vec2.ToString())
    }
	vec2.AddWithinBounds(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: 3, Y: 3})
	if !vec2.Equal(ZeroVector2[int]()) {
		t.Error(
			"vec2 should be Vector2(1, 1) + Vector2(3, 3) => in bounds of Vector2(2, 2) = Vector2(0, 0) and is",
			vec2.ToString(),
		)
	}
	vec2.AddWithinBounds(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: -3, Y: -3})
	if !vec2.Equal(Vector2[int]{X: -1, Y: -1}) {
		t.Error(
			"vec2 should be Vector2(0, 0) + Vector2(-3, -3) => in bounds of Vector2(2, 2) = Vector2(-1, -1) and is",
			vec2.ToString(),
		)
	}
	vec2.AddWithinBoundsAndForceAbsolute(Vector2[int]{X: 2, Y: 2}, Vector2[int]{X: -2, Y: -2})
	if !vec2.Equal(OneVector2[int]()) {
		t.Error(
			"vec2 should be Vector2(-1, -1) + Vector2(-2, -2) => in bounds of Vector2(2, 2) => Abs() = Vector2(1, 1) and is",
			vec2.ToString(),
		)
	}
}

func TestVector3(t *testing.T) {
	vec3 := Vector3[int]{X: 1, Y: 1, Z: 1}
	if vec3.X != 1 || vec3.Y != 1 || vec3.Z != 1 {
		t.Error("vec3 should be Vector3(1, 1, 1) and is", vec3.ToString())
	}
	if !vec3.Equal(OneVector3[int]()) {
		t.Error("vec3 should be Vector3(1, 1, 1) and is", vec3.ToString())
	}
	vec3.Add(OneVector3[int]())
	if !vec3.Equal(Vector3[int]{X: 2, Y: 2, Z: 2}){
        t.Error("vec3 should be Vector3(1, 1, 1) + Vector3(1, 1, 1) = Vector3(2, 2, 2) and is", vec3.ToString())
    }
	vec3.Substract(OneVector3[int]())
	if !vec3.Equal(OneVector3[int]()) {
        t.Error("vec3 should be Vector3(2, 2, 2) - Vector3(1, 1, 1) = Vector3(1, 1, 1) and is", vec3.ToString())
    }
	vec3.Multiply(Vector3[int]{X: 2, Y: 2, Z: 2})
	if !vec3.Equal(Vector3[int]{X: 2, Y: 2, Z: 2}) {
        t.Error("vec3 should be Vector3(1, 1, 1) * Vector3(2, 2, 2) = Vector3(2, 2, 2) and is", vec3.ToString())
    }
	vec3.Divide(Vector3[int]{X: 2, Y: 2, Z: 2})
	if !vec3.Equal(OneVector3[int]()) {
        t.Error("vec3 should be Vector3(2, 2, 2) / Vector3(2, 2, 2) = Vector3(1, 1, 1) and is", vec3.ToString())
    }
	vec3.AddWithinBounds(Vector3[int]{X: 2, Y: 2, Z: 2}, Vector3[int]{X: 3, Y: 3, Z: 3})
	if !vec3.Equal(ZeroVector3[int]()) {
		t.Error(
			"vec3 should be Vector3(1, 1, 1) + Vector3(3, 3, 3) => in bounds of Vector3(2, 2, 2) = Vector3(0, 0, 0) and is",
			vec3.ToString(),
		)
	}
	vec3.AddWithinBounds(Vector3[int]{X: 2, Y: 2, Z: 2}, Vector3[int]{X: -3, Y: -3, Z: -3})
	if !vec3.Equal(Vector3[int]{X: -1, Y: -1, Z: -1}) {
		t.Error(
			"vec3 should be Vector3(0, 0, 0) + Vector3(-3, -3, -3) => in bounds of Vector3(2, 2, 2) = Vector3(-1, -1, -1) and is",
			vec3.ToString(),
		)
	}
	vec3.AddWithinBoundsAndForceAbsolute(Vector3[int]{X: 2, Y: 2, Z: 2}, Vector3[int]{X: -2, Y: -2, Z: -2})
	if !vec3.Equal(OneVector3[int]()) {
		t.Error(
			"vec3 should be Vector3(-1, -1, -1) + Vector3(-2, -2, -2) => in bounds of Vector3(2, 2, 2) => Abs() = Vector3(1, 1, 1) and is",
			vec3.ToString(),
		)
	}
}

func TestVector4(t *testing.T) {
	vec4 := Vector4[int]{X: 1, Y: 1, Z: 1, W: 1}
	if vec4.X != 1 || vec4.Y != 1 || vec4.Z != 1 {
		t.Error("vec4 should be Vector4(1, 1, 1, 1) and is", vec4.ToString())
	}
	if !vec4.Equal(OneVector4[int]()) {
		t.Error("vec4 should be Vector4(1, 1, 1, 1) and is", vec4.ToString())
	}
	vec4.Add(OneVector4[int]())
	if !vec4.Equal(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2}){
        t.Error("vec4 should be Vector4(1, 1, 1, 1) + Vector4(1, 1, 1, 1) = Vector4(2, 2, 2, 2) and is", vec4.ToString())
    }
	vec4.Substract(OneVector4[int]())
	if !vec4.Equal(OneVector4[int]()) {
        t.Error("vec4 should be Vector4(2, 2, 2, 2) - Vector4(1, 1, 1, 1) = Vector4(1, 1, 1, 1) and is", vec4.ToString())
    }
	vec4.Multiply(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2})
	if !vec4.Equal(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2}) {
        t.Error("vec4 should be Vector4(1, 1, 1, 1) * Vector4(2, 2, 2, 2) = Vector4(2, 2, 2, 2) and is", vec4.ToString())
    }
	vec4.Divide(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2})
	if !vec4.Equal(OneVector4[int]()) {
        t.Error("vec4 should be Vector4(2, 2, 2, 2) / Vector4(2, 2, 2, 2) = Vector4(1, 1, 1, 1) and is", vec4.ToString())
    }
	vec4.AddWithinBounds(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2}, Vector4[int]{X: 3, Y: 3, Z: 3, W: 3})
	if !vec4.Equal(ZeroVector4[int]()) {
		t.Error(
			"vec4 should be Vector4(1, 1, 1, 1) + Vector4(3, 3, 3, 3) => in bounds of Vector4(2, 2, 2, 2) = Vector4(0, 0, 0, 0) and is",
			vec4.ToString(),
		)
	}
	vec4.AddWithinBounds(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2}, Vector4[int]{X: -3, Y: -3, Z: -3, W: -3})
	if !vec4.Equal(Vector4[int]{X: -1, Y: -1, Z: -1, W: -1}) {
		t.Error(
			"vec4 should be Vector4(0, 0, 0, 0) + Vector4(-3, -3, -3, -3) => in bounds of Vector4(2, 2, 2, 2) = Vector4(-1, -1, -1, -1) and is",
			vec4.ToString(),
		)
	}
	vec4.AddWithinBoundsAndForceAbsolute(Vector4[int]{X: 2, Y: 2, Z: 2, W: 2}, Vector4[int]{X: -2, Y: -2, Z: -2, W: -2})
	if !vec4.Equal(OneVector4[int]()) {
		t.Error(
			"vec4 should be Vector4(-1, -1, -1, -1) + Vector4(-2, -2, -2, -2) => in bounds of Vector4(2, 2, 2, 2) => Abs() = Vector4(1, 1, 1, 1) and is",
			vec4.ToString(),
		)
	}
}