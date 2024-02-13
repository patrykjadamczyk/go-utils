package math

import (
	"golang.org/x/exp/constraints"
)

// Clamp clamps number within the inclusive lower and upper bounds.
func Clamp[T constraints.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
