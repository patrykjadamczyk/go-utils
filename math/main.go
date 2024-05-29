package math

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var v T
		return v
	}

	max := items[0]

	for _, item := range items {
		if item > max {
			max = item
		}
	}

	return max
}

func Min[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var v T
		return v
	}

	min := items[0]

	for _, item := range items {
		if item < min {
			min = item
		}
	}

	return min
}

func abs[T ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~int](n T) T {
	if n >= 0 {
		return n
	}
	if n < 0 {
		return n * T(-1)
	}
	return n
}

func Abs[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint](
	n T,
) T {
	var v T
	value := any(v)
	switch value.(type) {
	case float64:
		return ensureType[T](abs(ensureType[float64](n)))
	case float32:
		return ensureType[T](abs(ensureType[float32](n)))
	case int:
		return ensureType[T](abs(ensureType[int](n)))
	case int8:
		return ensureType[T](abs(ensureType[int8](n)))
	case int16:
		return ensureType[T](abs(ensureType[int16](n)))
	case int32:
		return ensureType[T](abs(ensureType[int32](n)))
	case int64:
		return ensureType[T](abs(ensureType[int64](n)))
	}
	return n
}

func NaN[T ~float32 | ~float64]() T {
	var v T
	return v / v
}

func IsNaN[T ~float32 | ~float64](n T) bool {
	return n != n
}
