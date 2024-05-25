package math

import (
	"fmt"
)

// Vector4D
type Vector4[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint] struct {
	X int
	Y int
	Z int
	W int
}

// Add other vector to current one
func (vec *Vector4[T]) Add(other Vector4[T]) {
    vec.X += other.X
    vec.Y += other.Y
    vec.Z += other.Z
    vec.W += other.W
}

// Substract other vector from current one
func (vec *Vector4[T]) Substract(other Vector4[T]) {
    vec.X -= other.X
    vec.Y -= other.Y
    vec.Z -= other.Z
    vec.W -= other.W
}

// Multiply other vector with current one
func (vec *Vector4[T]) Multiply(other Vector4[T]) {
    vec.X *= other.X
    vec.Y *= other.Y
    vec.Z *= other.Z
    vec.W *= other.W
}

// Divide current vector by other one
func (vec *Vector4[T]) Divide(other Vector4[T]) {
    vec.X /= other.X
    vec.Y /= other.Y
    vec.Z /= other.Z
    vec.W /= other.W
}

// AddWithinBounds Add update vector to current one within bounds of bounds vector
func (vec *Vector4[T]) AddWithinBounds(bounds Vector4[T], update Vector4[T]) {
	vec.X = (vec.X + update.X) % bounds.X
	vec.Y = (vec.Y + update.Y) % bounds.Y
	vec.Z = (vec.Z + update.Z) % bounds.Z
	vec.W = (vec.W + update.W) % bounds.W
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector4[T]) AddWithinBoundsAndForceAbsolute(bounds Vector4[T], update Vector4[T]) {
	vec.X = Abs((vec.X + update.X) % bounds.X)
	vec.Y = Abs((vec.Y + update.Y) % bounds.Y)
	vec.Z = Abs((vec.Z + update.Z) % bounds.Z)
	vec.W = Abs((vec.W + update.W) % bounds.W)
}

// Equality check with other vector
func (vec *Vector4[T]) Equal(other Vector4[T]) bool {
	return vec.X == other.X && vec.Y == other.Y && vec.Z == other.Z && vec.W == other.W
}

// String representation of vector
func (vec Vector4[T]) ToString() string {
	return fmt.Sprintf("Vector4(%v,%v,%v,%v)", vec.X, vec.Y, vec.Z, vec.W)
}

// Zero Vector2D
func ZeroVector4[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint]() Vector4[T] {
	return Vector4[T]{X: 0, Y: 0, Z: 0, W: 0}
}

// One Vector2D
func OneVector4[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint]() Vector4[T] {
	return Vector4[T]{X: 1, Y: 1, Z: 1, W: 1}
}
