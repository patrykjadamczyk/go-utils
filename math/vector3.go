package math

import (
	"fmt"
)

// Vector3D
type Vector3[T VectorUnderlyingType] struct {
	X T
	Y T
	Z T
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
	vec.X = Mod((vec.X + update.X), bounds.X)
	vec.Y = Mod((vec.Y + update.Y), bounds.Y)
	vec.Z = Mod((vec.Z + update.Z), bounds.Z)
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector3[T]) AddWithinBoundsAndForceAbsolute(bounds Vector3[T], update Vector3[T]) {
	vec.X = Abs(Mod((vec.X + update.X), bounds.X))
	vec.Y = Abs(Mod((vec.Y + update.Y), bounds.Y))
	vec.Z = Abs(Mod((vec.Z + update.Z), bounds.Z))
}

// Equality check with other vector
func (vec *Vector3[T]) Equal(other Vector3[T]) bool {
	return vec.X == other.X && vec.Y == other.Y && vec.Z == other.Z
}

// String representation of vector
func (vec Vector3[T]) ToString() string {
	return fmt.Sprintf("Vector3(%v,%v,%v)", vec.X, vec.Y, vec.Z)
}

// Zero Vector3D
func ZeroVector3[T VectorUnderlyingType]() Vector3[T] {
	return Vector3[T]{X: 0, Y: 0, Z: 0}
}

// One Vector3D
func OneVector3[T VectorUnderlyingType]() Vector3[T] {
	return Vector3[T]{X: 1, Y: 1, Z: 1}
}

// Make Vector3D
func MakeVector3[T VectorUnderlyingType](x, y, z T) Vector3[T] {
	return Vector3[T]{X: x, Y: y, Z: z}
}

// Make Same Vector3D where all values of the vector are the same number specified
func SameVector3[T VectorUnderlyingType](number T) Vector3[T] {
	return MakeVector3(number, number, number)
}
