package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

type testMutStruct struct {
	test int
}

func TestMut(t *testing.T) {
	tv := MutRefValue(testMutStruct{test: 1})
	if tv.Get().test != 1 {
		t.Error("MutableRefValue doesn't hold value correctly")
	}
	tf := func (mv MutableRefValue[testMutStruct]) {
        mv.Set(testMutStruct{test: 2})
    }
    tf(tv)
	if tv.Get().test != 2 {
		t.Error("MutableRefValue doesn't change in function correctly")
	}
}
