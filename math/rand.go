package math

import (
	"math"
	"math/rand/v2"
)

func getRandomNumber[
	T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint | ~complex64 | ~complex128,
]() any {
	var v1 T
	var v2 any = v1
	switch v2.(type) {
	case int8:
		return int8(math.Floor(math.MaxInt8 * 2.0 * rand.Float64()))
	case int16:
		return int16(math.Floor(math.MaxInt16 * 2.0 * rand.Float64()))
	case int32:
		return int32(math.Floor(math.MaxInt32 * 2.0 * rand.Float64()))
	case int64:
		return int64(math.Floor(math.MaxInt64 * 2.0 * rand.Float64()))
	case int:
		return int(math.Floor(math.MaxInt * 2.0 * rand.Float64()))
	case uint8:
		return uint8(rand.Uint32N(math.MaxInt8))
	case uint16:
		return uint16(rand.Uint32N(math.MaxInt16))
	case uint32:
		return uint32(rand.Uint32N(math.MaxInt32))
	case uint64:
		return uint64(rand.Uint64N(math.MaxInt64))
	case uint:
		return uint(rand.Uint64N(math.MaxInt))
	case float32:
		if rand.Uint32N(2)%2 == 0 {
			return rand.Float32() * math.MaxFloat32
		}
		return -1 * rand.Float32() * math.MaxFloat32
	case float64:
		if rand.Uint32N(2)%2 == 0 {
			return rand.Float64() * math.MaxFloat64
		}
		return -1 * rand.Float64() * math.MaxFloat64
	case complex64:
		s1 := float32(rand.Uint32N(2) % 2)
		s2 := float32(rand.Uint32N(2) % 2)
		r := (s1 * -1.0) * rand.Float32() * math.MaxFloat32
		i := (s2 * -1.0) * rand.Float32() * math.MaxFloat32
		return complex(r, i)
	case complex128:
		s1 := float64(rand.Uint32N(2) % 2)
		s2 := float64(rand.Uint32N(2) % 2)
		r := (s1 * -1.0) * rand.Float64() * math.MaxFloat64
		i := (s2 * -1.0) * rand.Float64() * math.MaxFloat64
		return complex(r, i)
	default:
		return 0
	}
}

func GetRandomNumber[
	T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint,
]() T {
	if v, ok := getRandomNumber[T]().(T); ok {
		return v
	}
	var v T
	return v
}
