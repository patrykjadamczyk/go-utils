package compareenum_test

import (
	"testing"

	compare "github.com/patrykjadamczyk/go-utils/base/CompareEnum"
	"github.com/patrykjadamczyk/go-utils/math"
)

func TestToCompareEnum(t *testing.T) {
	test := func () {
		// int
		/// int8
		t1i8, t2i8 := math.GetRandomNumber[int8](), math.GetRandomNumber[int8]()

		if t1i8 > t2i8 && compare.ToCompareEnum(t1i8, t2i8) != compare.Greater {
			t.Error("Comparison of int8 > int8 failed")
		}
		if t1i8 == t2i8 && compare.ToCompareEnum(t1i8, t2i8) != compare.Equal {
			t.Error("Comparison of int8 == int8 failed")
		}
		if t1i8 < t2i8 && compare.ToCompareEnum(t1i8, t2i8) != compare.Less {
			t.Error("Comparison of int8 < int8 failed")
		}

		/// int16
		t1i16, t2i16 := math.GetRandomNumber[int16](), math.GetRandomNumber[int16]()

		if t1i16 > t2i16 && compare.ToCompareEnum(t1i16, t2i16) != compare.Greater {
			t.Error("Comparison of int16 > int16 failed")
		}
		if t1i16 == t2i16 && compare.ToCompareEnum(t1i16, t2i16) != compare.Equal {
			t.Error("Comparison of int16 == int16 failed")
		}
		if t1i16 < t2i16 && compare.ToCompareEnum(t1i16, t2i16) != compare.Less {
			t.Error("Comparison of int16 < int16 failed")
		}

		/// int32
		t1i32, t2i32 := math.GetRandomNumber[int32](), math.GetRandomNumber[int32]()

		if t1i32 > t2i32 && compare.ToCompareEnum(t1i32, t2i32) != compare.Greater {
			t.Error("Comparison of int32 > int32 failed")
		}
		if t1i32 == t2i32 && compare.ToCompareEnum(t1i32, t2i32) != compare.Equal {
			t.Error("Comparison of int32 == int32 failed")
		}
		if t1i32 < t2i32 && compare.ToCompareEnum(t1i32, t2i32) != compare.Less {
			t.Error("Comparison of int32 < int32 failed")
		}

		/// int64
		t1i64, t2i64 := math.GetRandomNumber[int64](), math.GetRandomNumber[int64]()

		if t1i64 > t2i64 && compare.ToCompareEnum(t1i64, t2i64) != compare.Greater {
			t.Error("Comparison of int64 > int64 failed")
		}
		if t1i64 == t2i64 && compare.ToCompareEnum(t1i64, t2i64) != compare.Equal {
			t.Error("Comparison of int64 == int64 failed")
		}
		if t1i64 < t2i64 && compare.ToCompareEnum(t1i64, t2i64) != compare.Less {
			t.Error("Comparison of int64 < int64 failed")
		}

		/// int
		t1i, t2i := math.GetRandomNumber[int](), math.GetRandomNumber[int]()

		if t1i > t2i && compare.ToCompareEnum(t1i, t2i) != compare.Greater {
			t.Error("Comparison of int > int failed")
		}
		if t1i == t2i && compare.ToCompareEnum(t1i, t2i) != compare.Equal {
			t.Error("Comparison of int == int failed")
		}
		if t1i < t2i && compare.ToCompareEnum(t1i, t2i) != compare.Less {
			t.Error("Comparison of int < int failed")
		}

		// uint
		/// uint8
		t1ui8, t2ui8 := math.GetRandomNumber[uint8](), math.GetRandomNumber[uint8]()

		if t1ui8 > t2ui8 && compare.ToCompareEnum(t1ui8, t2ui8) != compare.Greater {
			t.Error("Comparison of uint8 > uint8 failed")
		}
		if t1ui8 == t2ui8 && compare.ToCompareEnum(t1ui8, t2ui8) != compare.Equal {
			t.Error("Comparison of uint8 == uint8 failed")
		}
		if t1ui8 < t2ui8 && compare.ToCompareEnum(t1ui8, t2ui8) != compare.Less {
			t.Error("Comparison of uint8 < uint8 failed")
		}

		/// uint16
		t1ui16, t2ui16 := math.GetRandomNumber[uint16](), math.GetRandomNumber[uint16]()

		if t1ui16 > t2ui16 && compare.ToCompareEnum(t1ui16, t2ui16) != compare.Greater {
			t.Error("Comparison of uint16 > uint16 failed")
		}
		if t1ui16 == t2ui16 && compare.ToCompareEnum(t1ui16, t2ui16) != compare.Equal {
			t.Error("Comparison of uint16 == uint16 failed")
		}
		if t1ui16 < t2ui16 && compare.ToCompareEnum(t1ui16, t2ui16) != compare.Less {
			t.Error("Comparison of uint16 < uint16 failed")
		}

		/// uint32
		t1ui32, t2ui32 := math.GetRandomNumber[uint32](), math.GetRandomNumber[uint32]()

		if t1ui32 > t2ui32 && compare.ToCompareEnum(t1ui32, t2ui32) != compare.Greater {
			t.Error("Comparison of uint32 > uint32 failed")
		}
		if t1ui32 == t2ui32 && compare.ToCompareEnum(t1ui32, t2ui32) != compare.Equal {
			t.Error("Comparison of uint32 == uint32 failed")
		}
		if t1ui32 < t2ui32 && compare.ToCompareEnum(t1ui32, t2ui32) != compare.Less {
			t.Error("Comparison of uint32 < uint32 failed")
		}

		/// uint64
		t1ui64, t2ui64 := math.GetRandomNumber[uint64](), math.GetRandomNumber[uint64]()

		if t1ui64 > t2ui64 && compare.ToCompareEnum(t1ui64, t2ui64) != compare.Greater {
			t.Error("Comparison of uint64 > uint64 failed")
		}
		if t1ui64 == t2ui64 && compare.ToCompareEnum(t1ui64, t2ui64) != compare.Equal {
			t.Error("Comparison of uint64 == uint64 failed")
		}
		if t1ui64 < t2ui64 && compare.ToCompareEnum(t1ui64, t2ui64) != compare.Less {
			t.Error("Comparison of uint64 < uint64 failed")
		}

		/// uint
		t1ui, t2ui := math.GetRandomNumber[uint](), math.GetRandomNumber[uint]()

		if t1ui > t2ui && compare.ToCompareEnum(t1ui, t2ui) != compare.Greater {
			t.Error("Comparison of uint > uint failed")
		}
		if t1ui == t2ui && compare.ToCompareEnum(t1ui, t2ui) != compare.Equal {
			t.Error("Comparison of uint == uint failed")
		}
		if t1ui < t2ui && compare.ToCompareEnum(t1ui, t2ui) != compare.Less {
			t.Error("Comparison of uint < uint failed")
		}

		// float
		/// float32
		t1f32, t2f32 := math.GetRandomNumber[float32](), math.GetRandomNumber[float32]()

		if t1f32 > t2f32 && compare.ToCompareEnum(t1f32, t2f32) != compare.Greater {
			t.Error("Comparison of float32 > float32 failed")
		}
		if t1f32 == t2f32 && compare.ToCompareEnum(t1f32, t2f32) != compare.Equal {
			t.Error("Comparison of float32 == float32 failed")
		}
		if t1f32 < t2f32 && compare.ToCompareEnum(t1f32, t2f32) != compare.Less {
			t.Error("Comparison of float32 < float32 failed")
		}

		/// float64
		t1f64, t2f64 := math.GetRandomNumber[float64](), math.GetRandomNumber[float64]()

		if t1f64 > t2f64 && compare.ToCompareEnum(t1f64, t2f64) != compare.Greater {
			t.Error("Comparison of float64 > float64 failed")
		}
		if t1f64 == t2f64 && compare.ToCompareEnum(t1f64, t2f64) != compare.Equal {
			t.Error("Comparison of float64 == float64 failed")
		}
		if t1f64 < t2f64 && compare.ToCompareEnum(t1f64, t2f64) != compare.Less {
			t.Error("Comparison of float64 < float64 failed")
		}

		// complex
		/// complex64
		t1c64, t2c64 := math.GetRandomNumber[complex64](), math.GetRandomNumber[complex64]()
		if real(t1c64) > real(t2c64) && imag(t1c64) > imag(t2c64) && compare.ComplexToCompareEnum(t1c64, t2c64) != compare.Greater {
			t.Error("Comparison of complex64 > complex64 failed")
		}
		if real(t1c64) == real(t2c64) && imag(t1c64) == imag(t2c64) && compare.ComplexToCompareEnum(t1c64, t2c64) != compare.Equal {
			t.Error("Comparison of complex64 == complex64 failed")
		}
		if real(t1c64) < real(t2c64) && imag(t1c64) < imag(t2c64) && compare.ComplexToCompareEnum(t1c64, t2c64) != compare.Less {
			t.Error("Comparison of complex64 < complex64 failed")
		}
		if real(t1c64) < real(t2c64) && imag(t1c64) >= imag(t2c64) && compare.ComplexToCompareEnum(t1c64, t2c64) != compare.Uncomparable {
			t.Error("Comparison of 2 complex64 which can't be correctly compared by just real and imag parts failed")
		}

		/// complex128
		t1c128, t2c128 := math.GetRandomNumber[complex128](), math.GetRandomNumber[complex128]()
		if real(t1c128) > real(t2c128) && imag(t1c128) > imag(t2c128) && compare.ComplexToCompareEnum(t1c128, t2c128) != compare.Greater {
			t.Error("Comparison of complex128 > complex128 failed")
		}
		if real(t1c128) == real(t2c128) && imag(t1c128) == imag(t2c128) && compare.ComplexToCompareEnum(t1c128, t2c128) != compare.Equal {
			t.Error("Comparison of complex128 == complex128 failed")
		}
		if real(t1c128) < real(t2c128) && imag(t1c128) < imag(t2c128) && compare.ComplexToCompareEnum(t1c128, t2c128) != compare.Less {
			t.Error("Comparison of complex128 < complex128 failed")
		}
		if real(t1c128) < real(t2c128) && imag(t1c128) >= imag(t2c128) && compare.ComplexToCompareEnum(t1c128, t2c128) != compare.Uncomparable {
			t.Error("Comparison of 2 complex128 which can't be correctly compared by just real and imag parts failed")
		}
	}

	// Run test 100 times for good check through random numbers
	for range make([]bool, 100) {
		test()
	}
}
