package tests

import (
	"testing"
)

func TestIfPanics(t *testing.T, f func(), msg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg)
		}
	}()
	f()
	t.Error(msg)
}

func TestIfNotPanics(t *testing.T, f func(), msg string) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(msg)
		}
	}()
	f()
}
