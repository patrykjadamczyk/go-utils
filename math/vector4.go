package math

import (
	"fmt"
)

// Vector4D
type Vector4[T VectorUnderlyingType] struct {
	X T
	Y T
	Z T
	W T
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
	vec.X = Mod((vec.X + update.X), bounds.X)
	vec.Y = Mod((vec.Y + update.Y), bounds.Y)
	vec.Z = Mod((vec.Z + update.Z), bounds.Z)
	vec.W = Mod((vec.W + update.W), bounds.W)
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector4[T]) AddWithinBoundsAndForceAbsolute(bounds Vector4[T], update Vector4[T]) {
	vec.X = Abs(Mod((vec.X + update.X), bounds.X))
	vec.Y = Abs(Mod((vec.Y + update.Y), bounds.Y))
	vec.Z = Abs(Mod((vec.Z + update.Z), bounds.Z))
	vec.W = Abs(Mod((vec.W + update.W), bounds.W))
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
func ZeroVector4[T VectorUnderlyingType]() Vector4[T] {
	return Vector4[T]{X: 0, Y: 0, Z: 0, W: 0}
}

// One Vector2D
func OneVector4[T VectorUnderlyingType]() Vector4[T] {
	return Vector4[T]{X: 1, Y: 1, Z: 1, W: 1}
}

// Make Vector4D
func MakeVector4[T VectorUnderlyingType](x, y, z, w T) Vector4[T] {
	return Vector4[T]{X: x, Y: y, Z: z}
}

// Make Same Vector4D where all values of the vector are the same number specified
func SameVector4[T VectorUnderlyingType](number T) Vector4[T] {
	return MakeVector4(number, number, number, number)
}
