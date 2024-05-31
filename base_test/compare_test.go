package base_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
	compare "github.com/patrykjadamczyk/go-utils/base/CompareEnum"
	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

type testCompareEqualityStruct struct {
}

func (t testCompareEqualityStruct) Equals(v any) bool {
    return true
}

type testCompareLessStruct struct {
}

func (t testCompareLessStruct) LessThan(v any) bool {
    return true
}

type testCompareGreaterStruct struct {
}

func (t testCompareGreaterStruct) GreaterThan(v any) bool {
    return true
}

func TestCompare(t *testing.T) {
	// Enforce that these structs implement interfaces
	_ = interfaces.IEqualsAny(testCompareEqualityStruct{})
	_ = interfaces.ILessThanAny(testCompareLessStruct{})
	_ = interfaces.IGreaterThanAny(testCompareGreaterStruct{})
	// Specify tests
	tests := [][][2]any{
		// ==
		{
			// intX
			{int8(1), int8(1)},
			{int16(1), int16(1)},
			{int32(1), int32(1)},
			{int64(1), int64(1)},
			{int8(1), int8(1)},
			{int8(1), int16(1)},
			{int8(1), int32(1)},
			{int8(1), int64(1)},
			{int16(1), int8(1)},
			{int16(1), int16(1)},
			{int16(1), int32(1)},
			{int16(1), int64(1)},
			{int32(1), int8(1)},
			{int32(1), int16(1)},
			{int32(1), int32(1)},
			{int32(1), int64(1)},
			{int64(1), int8(1)},
			{int64(1), int16(1)},
			{int64(1), int32(1)},
			{int64(1), int64(1)},
			// int
			{int(1), int(1)},
			// uintX
			{uint8(1), uint8(1)},
			{uint16(1), uint16(1)},
			{uint32(1), uint32(1)},
			{uint64(1), uint64(1)},
			{uint8(1), uint8(1)},
			{uint8(1), uint16(1)},
			{uint8(1), uint32(1)},
			{uint8(1), uint64(1)},
			{uint16(1), uint8(1)},
			{uint16(1), uint16(1)},
			{uint16(1), uint32(1)},
			{uint16(1), uint64(1)},
			{uint32(1), uint8(1)},
			{uint32(1), uint16(1)},
			{uint32(1), uint32(1)},
			{uint32(1), uint64(1)},
			{uint64(1), uint8(1)},
			{uint64(1), uint16(1)},
			{uint64(1), uint32(1)},
			{uint64(1), uint64(1)},
			// uint
			{uint(1), uint(1)},
			// float
			{float32(1), float32(1)},
			{float64(1), float32(1)},
			{float32(1), float64(1)},
			{float64(1), float64(1)},
			// complex
			{complex64(complex(1, 1)), complex64(complex(1, 1))},
			{complex128(complex(1, 1)), complex64(complex(1, 1))},
			{complex64(complex(1, 1)), complex128(complex(1, 1))},
			{complex128(complex(1, 1)), complex128(complex(1, 1))},

			// bool
			{true, true},
			{false, false},

			// string
			{"", ""},

			// Structs
			{nil, nil},
			{testCompareEqualityStruct{}, testCompareEqualityStruct{}},
			{testCompareEqualityStruct{}, int(1)},
			{Null[int](), Null[int]()},
			{NullableValue(1), NullableValue(1)},
			{Ok(1), Ok(1)},
		},
		// >
		{
			// intX
			{int8(2), int8(1)},
			{int16(2), int16(1)},
			{int32(2), int32(1)},
			{int64(2), int64(1)},
			{int8(2), int8(1)},
			{int8(2), int16(1)},
			{int8(2), int32(1)},
			{int8(2), int64(1)},
			{int16(2), int8(1)},
			{int16(2), int16(1)},
			{int16(2), int32(1)},
			{int16(2), int64(1)},
			{int32(2), int8(1)},
			{int32(2), int16(1)},
			{int32(2), int32(1)},
			{int32(2), int64(1)},
			{int64(2), int8(1)},
			{int64(2), int16(1)},
			{int64(2), int32(1)},
			{int64(2), int64(1)},
			// int
			{int(2), int(1)},
			// uintX
			{uint8(2), uint8(1)},
			{uint16(2), uint16(1)},
			{uint32(2), uint32(1)},
			{uint64(2), uint64(1)},
			{uint8(2), uint8(1)},
			{uint8(2), uint16(1)},
			{uint8(2), uint32(1)},
			{uint8(2), uint64(1)},
			{uint16(2), uint8(1)},
			{uint16(2), uint16(1)},
			{uint16(2), uint32(1)},
			{uint16(2), uint64(1)},
			{uint32(2), uint8(1)},
			{uint32(2), uint16(1)},
			{uint32(2), uint32(1)},
			{uint32(2), uint64(1)},
			{uint64(2), uint8(1)},
			{uint64(2), uint16(1)},
			{uint64(2), uint32(1)},
			{uint64(2), uint64(1)},
			// uint
			{uint(2), uint(1)},
			// float
			{float32(2), float32(1)},
			{float64(2), float32(1)},
			{float32(2), float64(1)},
			{float64(2), float64(1)},
			// complex
			{complex64(complex(2, 2)), complex64(complex(1, 1))},
			{complex128(complex(2, 2)), complex64(complex(1, 1))},
			{complex64(complex(2, 2)), complex128(complex(1, 1))},
			{complex128(complex(2, 2)), complex128(complex(1, 1))},

			// Structs
			{int(1), nil},
			{testCompareGreaterStruct{}, testCompareGreaterStruct{}},
			{testCompareGreaterStruct{}, int(1)},
			{NullableValue(1), Null[int]()},
			{NullableValue(2), NullableValue(1)},
			{Ok(2), Ok(1)},
		},
		// <
		{
			// intX
			{int8(2), int8(3)},
			{int16(2), int16(3)},
			{int32(2), int32(3)},
			{int64(2), int64(3)},
			{int8(2), int8(3)},
			{int8(2), int16(3)},
			{int8(2), int32(3)},
			{int8(2), int64(3)},
			{int16(2), int8(3)},
			{int16(2), int16(3)},
			{int16(2), int32(3)},
			{int16(2), int64(3)},
			{int32(2), int8(3)},
			{int32(2), int16(3)},
			{int32(2), int32(3)},
			{int32(2), int64(3)},
			{int64(2), int8(3)},
			{int64(2), int16(3)},
			{int64(2), int32(3)},
			{int64(2), int64(3)},
			// int
			{int(2), int(3)},
			// uintX
			{uint8(2), uint8(3)},
			{uint16(2), uint16(3)},
			{uint32(2), uint32(3)},
			{uint64(2), uint64(3)},
			{uint8(2), uint8(3)},
			{uint8(2), uint16(3)},
			{uint8(2), uint32(3)},
			{uint8(2), uint64(3)},
			{uint16(2), uint8(3)},
			{uint16(2), uint16(3)},
			{uint16(2), uint32(3)},
			{uint16(2), uint64(3)},
			{uint32(2), uint8(3)},
			{uint32(2), uint16(3)},
			{uint32(2), uint32(3)},
			{uint32(2), uint64(3)},
			{uint64(2), uint8(3)},
			{uint64(2), uint16(3)},
			{uint64(2), uint32(3)},
			{uint64(2), uint64(3)},
			// uint
			{uint(2), uint(3)},
			// float
			{float32(2), float32(3)},
			{float64(2), float32(3)},
			{float32(2), float64(3)},
			{float64(2), float64(3)},
			// complex
			{complex64(complex(2, 2)), complex64(complex(3, 3))},
			{complex128(complex(2, 2)), complex64(complex(3, 3))},
			{complex64(complex(2, 2)), complex128(complex(3, 3))},
			{complex128(complex(2, 2)), complex128(complex(3, 3))},

			// Structs
			{nil, int(1)},
			{testCompareLessStruct{}, testCompareLessStruct{}},
			{testCompareLessStruct{}, int(1)},
			{NullableValue(-1), Null[int]()},
			{NullableValue(1), NullableValue(2)},
			{Ok(1), Ok(2)},
		},
		// Uncomparable
		{
			// bool
			{true, false},
			// string
			{"", ">"},
			// error
			{NewError("test"), NewError("test2")},
		},
	}
	translate := map[compare.CompareEnum]string{
		compare.None: "None",
		compare.Less: "Less",
		compare.Equal: "Equal",
		compare.Greater: "Greater",
		compare.Error: "Error",
		compare.Uncomparable: "Uncomparable",
	}

	for testTypeIndex, testType := range tests {
		for _, test := range testType {
			switch testTypeIndex {
			case 0:
				if uc := UniversalCompare(test[0], test[1]); uc != compare.Equal {
					t.Errorf(
						"Testing (%T)%#v == (%T)%#v failed, got %v",
						test[0],
						test[0],
						test[1],
						test[1],
						translate[uc],
					)
				}
			case 1:
				if uc := UniversalCompare(test[0], test[1]); uc != compare.Greater {
					t.Errorf(
						"Testing (%T)%#v > (%T)%#v failed, got %v",
						test[0],
						test[0],
						test[1],
						test[1],
						translate[uc],
					)
				}
			case 2:
				if uc := UniversalCompare(test[0], test[1]); uc != compare.Less {
					t.Errorf(
						"Testing (%T)%#v < (%T)%#v failed, got %v",
						test[0],
						test[0],
						test[1],
						test[1],
						translate[uc],
					)
				}
			case 3:
				if uc := UniversalCompare(test[0], test[1]); uc != compare.Uncomparable {
					t.Errorf(
						"Testing (%T)%#v and (%T)%#v are uncomparable failed, got %v",
						test[0],
						test[0],
						test[1],
						test[1],
						translate[uc],
					)
				}
			}
		}
	}
}
