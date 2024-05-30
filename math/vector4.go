package math

import (
	"fmt"
	"strconv"
	"strings"
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

// EqualWithinTolerance check if current vector is within tolerance vector of other vector
func (vec *Vector4[T]) EqualWithinTolerance(other Vector4[T], tolerance Vector4[T]) bool {
	ret := vec.X >= other.X-tolerance.X && vec.X <= other.X+tolerance.X
	ret = ret && vec.Y >= other.Y-tolerance.Y && vec.Y <= other.Y+tolerance.Y
	ret = ret && vec.Z >= other.Z-tolerance.Z && vec.Z <= other.Z+tolerance.Z
	ret = ret && vec.W >= other.W-tolerance.W && vec.W <= other.W+tolerance.W
	return ret
}

// String representation of vector
func (vec Vector4[T]) ToString() string {
	return fmt.Sprintf("Vector4(%v,%v,%v,%v)", vec.X, vec.Y, vec.Z, vec.W)
}

// String representation of vector
func (vec Vector4[T]) String() string {
	return vec.ToString()
}

// Change values of current vector to values from string representation
func (vec *Vector4[T]) FromString(v string) {
	vs := strings.TrimLeft(v, "Vector4")
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

	if len(vsat) != 4 {
		panic("AssertionError: Vector string malformed")
	}
	x := ensureType[T](vsat[0])
	y := ensureType[T](vsat[1])
	z := ensureType[T](vsat[2])
	w := ensureType[T](vsat[3])
	vec.X = x
	vec.Y = y
	vec.Z = z
	vec.W = w
}

// ContainsNaN checks if vector contains NaN
func (vec Vector4[T]) ContainsNaN() bool {
	return vec.X != vec.X || vec.Y != vec.Y || vec.Z != vec.Z || vec.W != vec.W
}

// ToTuple get tuple of cordinates of vector
func (vec Vector4[T]) ToTuple() (x, y, z, w T) {
	x = vec.X
	y = vec.Y
	z = vec.Z
	w = vec.W
	return
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
	return Vector4[T]{X: x, Y: y, Z: z, W: w}
}

// Make Same Vector4D where all values of the vector are the same number specified
func SameVector4[T VectorUnderlyingType](number T) Vector4[T] {
	return MakeVector4(number, number, number, number)
}
