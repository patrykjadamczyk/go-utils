package math

import "github.com/patrykjadamczyk/go-utils/base"

// Bitwise and
func BitwiseAnd[T base.AnyInt](a, b T) T {
	return a & b
}

// Bitwise or
func BitwiseOr[T base.AnyInt](a, b T) T {
	return a | b
}

// Bitwise xor
func BitwiseXor[T base.AnyInt](a, b T) T {
	return a ^ b
}

// Bitwise not
func BitwiseNot[T base.AnyInt](a T) T {
	return ^a
}

// Bitwise Left Shift
func BitwiseLeftShift[T base.AnyInt, ST int|uint](a T, shift ST) T {
	return a << shift
}

// Bitwise Right Shift
func BitwiseRightShift[T base.AnyInt, ST int|uint](a T, shift ST) T {
	return a >> shift
}
