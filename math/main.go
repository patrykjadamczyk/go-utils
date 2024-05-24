package math

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		return Null[T]().ValueOrZero()
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
		return Null[T]().ValueOrZero()
	}

	min := items[0]

	for _, item := range items {
		if item < min {
			min = item
		}
	}

	return min
}

func Abs[T SignedNumericType](n T) T {
	if n >= 0 {
		return n
	}
	if n < 0 {
		return n * T(-1)
	}
	return n
}
