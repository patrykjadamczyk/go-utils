package base_test

import (
	"testing"
	. "github.com/patrykjadamczyk/go-utils/base"
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
