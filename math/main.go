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

func Abs[T ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~int](n T) T {
	if n >= 0 {
		return n
	}
	if n < 0 {
		return n * T(-1)
	}
	return n
}

func NaN[T ~float32 | ~float64]() T {
	var v T
	return v/v
}

func IsNaN[T ~float32 | ~float64](n T) bool {
    return n != n
}
