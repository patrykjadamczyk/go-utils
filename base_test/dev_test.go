package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

type TestStruct struct {
}

func (t TestStruct) Dev() Result[string] {
	return EnsureObjectImplementsInterface[TestStruct, EnsuredObject](t)
}

func TestDevFunction(t *testing.T) {
	t1 := TestStruct{}
	t1d := t1.Dev()
	if t1d.IsError() {
		t.Error("Result 1 should not be error", t1d.DataValue, t1d.ErrorValue)
	}
	if t1d.UnwrapOr("error") != "OK" {
		t.Error("Result 1 should be OK", t1d.DataValue, t1d.ErrorValue)
	}
}
