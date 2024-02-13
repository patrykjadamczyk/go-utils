package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestEnsureType1(t *testing.T) {
	if EnsureType[int](any(1)) != 1 {
		t.Error("EnsureType should work on correct data")
	}
}

func TestEnsureType2(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("EnsureType should panic on invalid data")
		}
	}()
	EnsureType[string](1)
	t.Error("EnsureType should panic on invalid data")
}

func TestEnsureType3(t *testing.T) {
	if EnsureTypeResult[int](any(1)).IsError() {
		t.Error("EnsureTypeResult should have ok value on correct data")
	}
}

func TestEnsureType4(t *testing.T) {
	if !EnsureTypeResult[string](1).IsError() {
		t.Error("EnsureTypeResult should have err value on invalid data")
	}
}
