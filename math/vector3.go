package math

import (
	"fmt"
	"strconv"
	"strings"
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

// EqualWithinTolerance check if current vector is within tolerance vector of other vector
func (vec *Vector3[T]) EqualWithinTolerance(other Vector3[T], tolerance Vector3[T]) bool {
	ret := vec.X >= other.X-tolerance.X && vec.X <= other.X+tolerance.X
	ret = ret && vec.Y >= other.Y-tolerance.Y && vec.Y <= other.Y+tolerance.Y
	ret = ret && vec.Z >= other.Z-tolerance.Z && vec.Z <= other.Z+tolerance.Z
	return ret
}

// String representation of vector
func (vec Vector3[T]) ToString() string {
	return fmt.Sprintf("Vector3(%v,%v,%v)", vec.X, vec.Y, vec.Z)
}

// String representation of vector
func (vec Vector3[T]) String() string {
	return vec.ToString()
}

// Change values of current vector to values from string representation
func (vec *Vector3[T]) FromString(v string) {
	vs := strings.TrimLeft(v, "Vector3")
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

	if len(vsat) != 3 {
		panic("AssertionError: Vector string malformed")
	}
	x := ensureType[T](vsat[0])
	y := ensureType[T](vsat[1])
	z := ensureType[T](vsat[2])
	vec.X = x
	vec.Y = y
	vec.Z = z
}

// ContainsNaN checks if vector contains NaN
func (vec Vector3[T]) ContainsNaN() bool {
	return vec.X != vec.X || vec.Y != vec.Y || vec.Z != vec.Z
}

// ToTuple get tuple of cordinates of vector
func (vec Vector3[T]) ToTuple() (x, y, z T) {
	x = vec.X
	y = vec.Y
	z = vec.Z
	return
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
