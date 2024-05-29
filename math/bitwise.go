package math

// Bitwise and
func BitwiseAnd[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint](a, b T) T {
	return a & b
}

// Bitwise or
func BitwiseOr[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint](a, b T) T {
	return a | b
}

// Bitwise xor
func BitwiseXor[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint](a, b T) T {
	return a ^ b
}

// Bitwise not
func BitwiseNot[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint](a T) T {
	return ^a
}

// Bitwise Left Shift
func BitwiseLeftShift[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint, ST int | uint](a T, shift ST) T {
	return a << shift
}

// Bitwise Right Shift
func BitwiseRightShift[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint, ST int | uint](a T, shift ST) T {
	return a >> shift
}
