package math

import (
	"fmt"
	"strconv"
	"strings"
)

// Vector2D
type Vector2[T VectorUnderlyingType] struct {
	X T
	Y T
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
	vec.X = Mod((vec.X + update.X), bounds.X)
	vec.Y = Mod((vec.Y + update.Y), bounds.Y)
}

// AddWithinBoundsAndForceAbsolute Add update vector to current one within bounds of bounds vector and force value of vector to be positive
func (vec *Vector2[T]) AddWithinBoundsAndForceAbsolute(bounds Vector2[T], update Vector2[T]) {
	vec.X = Abs(Mod((vec.X + update.X), bounds.X))
	vec.Y = Abs(Mod((vec.Y + update.Y), bounds.Y))
}

// Equality check with other vector
func (vec *Vector2[T]) Equal(other Vector2[T]) bool {
	return vec.X == other.X && vec.Y == other.Y
}

// EqualWithinTolerance check if current vector is within tolerance vector of other vector
func (vec *Vector2[T]) EqualWithinTolerance(other Vector2[T], tolerance Vector2[T]) bool {
	ret := vec.X >= other.X-tolerance.X && vec.X <= other.X+tolerance.X
	ret = ret && vec.Y >= other.Y-tolerance.Y && vec.Y <= other.Y+tolerance.Y
	return ret
}

// String representation of vector
func (vec Vector2[T]) ToString() string {
	return fmt.Sprintf("Vector2(%v,%v)", vec.X, vec.Y)
}

// String representation of vector
func (vec Vector2[T]) String() string {
	return vec.ToString()
}

// Change values of current vector to values from string representation
func (vec *Vector2[T]) FromString(v string) {
	vs := strings.TrimLeft(v, "Vector2")
	vs = strings.TrimLeft(vs, "(")
	vs = strings.TrimRight(vs, ")")
	vsa := strings.Split(vs, ",")
	var vsat = make([]any, 0)
	var _typeswitch T
	typeswitch := any(_typeswitch)
	for _, vsi := range vsa {
		switch typeswitch.(type) {
		case int:
			a, _ := strconv.ParseInt(vsi, 10, 64)
			vsat = append(vsat, int(a))
		case int8:
			a, _ := strconv.ParseInt(vsi, 10, 8)
			vsat = append(vsat, int8(a))
		case int16:
			a, _ := strconv.ParseInt(vsi, 10, 16)
			vsat = append(vsat, int16(a))
		case int32:
			a, _ := strconv.ParseInt(vsi, 10, 32)
			vsat = append(vsat, int32(a))
		case int64:
			a, _ := strconv.ParseInt(vsi, 10, 64)
			vsat = append(vsat, int64(a))
		case uint:
			a, _ := strconv.ParseUint(vsi, 10, 64)
			vsat = append(vsat, uint(a))
		case uint8:
			a, _ := strconv.ParseUint(vsi, 10, 8)
			vsat = append(vsat, uint8(a))
		case uint16:
			a, _ := strconv.ParseUint(vsi, 10, 16)
			vsat = append(vsat, uint16(a))
		case uint32:
			a, _ := strconv.ParseUint(vsi, 10, 32)
			vsat = append(vsat, uint32(a))
		case uint64:
			a, _ := strconv.ParseUint(vsi, 10, 64)
			vsat = append(vsat, uint64(a))
		case float32:
			a, _ := strconv.ParseFloat(vsi, 32)
			vsat = append(vsat, float32(a))
		case float64:
			a, _ := strconv.ParseFloat(vsi, 64)
			vsat = append(vsat, float64(a))
		}
	}

	if len(vsat) != 2 {
		panic("AssertionError: Vector string malformed")
	}
	x := ensureType[T](vsat[0])
	y := ensureType[T](vsat[1])
	vec.X = x
	vec.Y = y
}

// ContainsNaN checks if vector contains NaN
func (vec Vector2[T]) ContainsNaN() bool {
	return vec.X != vec.X || vec.Y != vec.Y
}

// ToTuple get tuple of cordinates of vector
func (vec Vector2[T]) ToTuple() (x, y T) {
	x = vec.X
	y = vec.Y
	return
}

// Zero Vector2D
func ZeroVector2[T VectorUnderlyingType]() Vector2[T] {
	return Vector2[T]{X: 0, Y: 0}
}

// One Vector2D
func OneVector2[T VectorUnderlyingType]() Vector2[T] {
	return Vector2[T]{X: 1, Y: 1}
}

// Make Vector2D
func MakeVector2[T VectorUnderlyingType](x, y T) Vector2[T] {
	return Vector2[T]{X: x, Y: y}
}

// Make Same Vector2D where all values of the vector are the same number specified
func SameVector2[T VectorUnderlyingType](number T) Vector2[T] {
	return MakeVector2(number, number)
}
