package math

import (
	"fmt"

	. "github.com/patrykjadamczyk/go-utils/base"
)

// Vector3D
type Vector3[T NumericType] struct {
	X int
	Y int
	Z int
}

// Add other vector to current one
func (vec *Vector3[T]) Add(other Vector3[T]) {
    vec.X += other.X
    vec.Y += other.Y
    vec.Z += other.Z
}

// Substract other vector from current one
func (vec *Vector3[T]) Substract(other Vector3[T]) {
    vec.X -= other.X
    vec.Y -= other.Y
    vec.Z -= other.Z
}

// Multiply other vector with current one
func (vec *Vector3[T]) Multiply(other Vector3[T]) {
    vec.X *= other.X
    vec.Y *= other.Y
    vec.Z *= other.Z
}

// Divide current vector by other one
func (vec *Vector3[T]) Divide(other Vector3[T]) {
    vec.X /= other.X
    vec.Y /= other.Y
    vec.Z /= other.Z
}

// AddWithinBounds Add update vector to current one within bounds of bounds vector
func (vec *Vector3[T]) AddWithinBounds(bounds Vector3[T], update Vector3[T]) {
	vec.X = (vec.X + update.X) % bounds.X
	vec.Y = (vec.Y + update.Y) % bounds.Y
	vec.Z = (vec.Z + update.Z) % bounds.Z
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector3[T]) AddWithinBoundsAndForceAbsolute(bounds Vector3[T], update Vector3[T]) {
	vec.X = Abs((vec.X + update.X) % bounds.X)
	vec.Y = Abs((vec.Y + update.Y) % bounds.Y)
	vec.Z = Abs((vec.Z + update.Z) % bounds.Z)
}

// Equality check with other vector
func (vec *Vector3[T]) Equal(other Vector3[T]) bool {
	return vec.X == other.X && vec.Y == other.Y && vec.Z == other.Z
}

// String representation of vector
func (vec Vector3[T]) ToString() string {
	return fmt.Sprintf("Vector3(%v,%v,%v)", vec.X, vec.Y, vec.Z)
}
