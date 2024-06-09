package CompareEnum

type CompareEnum uint8

const (
	// None - NaN
	None CompareEnum = iota
	// Less than v1 < v2
	Less
	// Equal v1 == v2
	Equal
	// Greater than v1 > v2
	Greater
	// Error in comparing
	Error
	// Uncomparable
	Uncomparable
)

func ToCompareEnum[
	T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint,
](
	v1 T,
	v2 T,
) CompareEnum {
	if v1 < v2 {
		return Less
	}
	if v1 == v2 {
		return Equal
	}
	if v1 > v2 {
		return Greater
	}
	return Error
}

func ComplexToCompareEnum(
	v1 any,
	v2 any,
) CompareEnum {
	var v1c complex128
	var v2c complex128
	test1 := false
	test2 := false
	if v1i, v1ok := v1.(complex64); v1ok {
		test1 = true
		v1c = complex(float64(real(v1i)), float64(imag(v1i)))
	}
	if v1i, v1ok := v1.(complex128); v1ok {
		test1 = true
		v1c = v1i
	}
	if v2i, v2ok := v2.(complex64); v2ok {
		test2 = true
		v2c = complex(float64(real(v2i)), float64(imag(v2i)))
	}
	if v2i, v2ok := v2.(complex128); v2ok {
		test2 = true
		v2c = v2i
	}

	if !test1 || !test2 {
		return Error
	}

	if real(v1c) < real(v2c) && imag(v1c) < imag(v2c) {
		return Less
	}
	if real(v1c) == real(v2c) && imag(v1c) == imag(v2c) {
		return Equal
	}
	if real(v1c) > real(v2c) && imag(v1c) > imag(v2c) {
		return Greater
	}
	return Uncomparable
}
