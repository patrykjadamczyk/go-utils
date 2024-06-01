package base_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

func TestUniversalEquals(t *testing.T) {
	// Enforce that these structs implement interfaces
	_ = interfaces.IEqualsAny(testCompareEqualityStruct{})
	// Specify tests
	tests := [][2]any{
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
	}

	for _, test := range tests {
		if uc := UniversalEquals(test[0], test[1]); !uc.GetValue() || uc.IsError() {
			t.Errorf(
				"Testing (%T)%#v == (%T)%#v failed, got Result(%v, %v)",
				test[0],
				test[0],
				test[1],
				test[1],
				uc.GetValue(),
				uc.IsError(),
			)
		}
	}
}
