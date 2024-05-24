package math

import (
	"fmt"
)

// Vector2D
type Vector2[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint] struct {
	X int
	Y int
}

// Add other vector to current one
func (vec *Vector2[T]) Add(other Vector2[T]) {
    vec.X += other.X
    vec.Y += other.Y
}

// Substract other vector from current one
func (vec *Vector2[T]) Substract(other Vector2[T]) {
    vec.X -= other.X
    vec.Y -= other.Y
}

// Multiply other vector with current one
func (vec *Vector2[T]) Multiply(other Vector2[T]) {
    vec.X *= other.X
    vec.Y *= other.Y
}

// Divide current vector by other one
func (vec *Vector2[T]) Divide(other Vector2[T]) {
    vec.X /= other.X
    vec.Y /= other.Y
}

// AddWithinBounds Add update vector to current one within bounds of bounds vector
func (vec *Vector2[T]) AddWithinBounds(bounds Vector2[T], update Vector2[T]) {
	vec.X = (vec.X + update.X) % bounds.X
	vec.Y = (vec.Y + update.Y) % bounds.Y
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector2[T]) AddWithinBoundsAndForceAbsolute(bounds Vector2[T], update Vector2[T]) {
	vec.X = Abs((vec.X + update.X) % bounds.X)
	vec.Y = Abs((vec.Y + update.Y) % bounds.Y)
}

// Equality check with other vector
func (vec *Vector2[T]) Equal(other Vector2[T]) bool {
	return vec.X == other.X && vec.Y == other.Y
}

// String representation of vector
func (vec Vector2[T]) ToString() string {
	return fmt.Sprintf("Vector2(%v,%v)", vec.X, vec.Y)
}
