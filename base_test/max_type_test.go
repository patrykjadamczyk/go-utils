package base_test

import (
	"fmt"
	. "github.com/patrykjadamczyk/go-utils/base"
	"math"
	"testing"
)

func TestMaxType(t *testing.T) {
	testCases := []struct {
		input    int64
		expected int64
	}{
		{input: int64(GetMaxForType[int]()), expected: math.MaxInt},
		{input: int64(GetMaxForType[int8]()), expected: math.MaxInt8},
		{input: int64(GetMaxForType[int16]()), expected: math.MaxInt16},
		{input: int64(GetMaxForType[int32]()), expected: math.MaxInt32},
		{input: int64(GetMaxForType[int64]()), expected: math.MaxInt64},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test MaxType i%d", i), func(t *testing.T) {
			if tc.input != tc.expected {
				t.Errorf("i%d: Expected %d, but got %d", i, tc.expected, tc.input)
			}
		})
	}
	testCases2 := []struct {
		input    uint64
		expected uint64
	}{
		{input: uint64(GetMaxForType[uint8]()), expected: uint64(math.MaxUint8)},
		{input: uint64(GetMaxForType[uint16]()), expected: uint64(math.MaxUint16)},
		{input: uint64(GetMaxForType[uint32]()), expected: uint64(math.MaxUint32)},
	}

	for i, tc := range testCases2 {
		t.Run(fmt.Sprintf("Test MaxType ui%d", i), func(t *testing.T) {
			if tc.input != tc.expected {
				t.Errorf("ui%d: Expected %d, but got %d", i, tc.expected, tc.input)
			}
		})
	}
	testCases3 := []struct {
		input    float64
		expected float64
	}{
		{input: float64(GetMaxForType[float32]()), expected: float64(math.MaxFloat32)},
		{input: GetMaxForType[float64](), expected: float64(math.MaxFloat64)},
	}

	for i, tc := range testCases3 {
		t.Run(fmt.Sprintf("Test MaxType f%d", i), func(t *testing.T) {
			if tc.input != tc.expected {
				t.Errorf("f%d: Expected %e, but got %e", i, tc.expected, tc.input)
			}
		})
	}
	test1i := GetMaxForType[uint]()
	t.Run("Test MaxType a1", func(t *testing.T) {
		if test1i != math.MaxUint {
			t.Errorf("a1: Expected math.MaxUint, but got %d", test1i)
		}
	})
	test2i := GetMaxForType[uint64]()
	t.Run("Test MaxType a2", func(t *testing.T) {
		if test2i != math.MaxUint64 {
			t.Errorf("a2: Expected math.MaxUint64, but got %d", test1i)
		}
	})
}

func TestMinType(t *testing.T) {
	testCases := []struct {
		input    int64
		expected int64
	}{
		{input: int64(GetMinForType[int]()), expected: math.MinInt},
		{input: int64(GetMinForType[int8]()), expected: math.MinInt8},
		{input: int64(GetMinForType[int16]()), expected: math.MinInt16},
		{input: int64(GetMinForType[int32]()), expected: math.MinInt32},
		{input: int64(GetMinForType[int64]()), expected: math.MinInt64},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test MinType i%d", i), func(t *testing.T) {
			if tc.input != tc.expected {
				t.Errorf("i%d: Expected %d, but got %d", i, tc.expected, tc.input)
			}
		})
	}
	testCases2 := []struct {
		input uint64
	}{
		{input: uint64(GetMinForType[uint]())},
		{input: uint64(GetMinForType[uint8]())},
		{input: uint64(GetMinForType[uint16]())},
		{input: uint64(GetMinForType[uint32]())},
		{input: uint64(GetMinForType[uint64]())},
	}

	for i, tc := range testCases2 {
		t.Run(fmt.Sprintf("Test MinType ui%d", i), func(t *testing.T) {
			if tc.input != 0 {
				t.Errorf("ui%d: Expected 0, but got %d", i, tc.input)
			}
		})
	}
	testCases3 := []struct {
		input    float64
		expected float64
	}{
		{input: float64(GetMinForType[float32]()), expected: -1 * float64(math.MaxFloat32)},
		{input: GetMinForType[float64](), expected: -1 * float64(math.MaxFloat64)},
	}

	for i, tc := range testCases3 {
		t.Run(fmt.Sprintf("Test MinType f%d", i), func(t *testing.T) {
			if tc.input != tc.expected {
				t.Errorf("f%d: Expected %e, but got %e", i, tc.expected, tc.input)
			}
		})
	}
}
