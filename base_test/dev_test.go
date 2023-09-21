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
	if t1.Dev().IsError() {
		t.Error("Result 1 should not be error")
	}
	if t1.Dev().UnwrapOr("error") != "OK" {
		t.Error("Result 1 should be OK")
	}
}
