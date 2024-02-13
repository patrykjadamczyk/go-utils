package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestDebugMode(t *testing.T) {
	debug := GetDebugMode()

	if debug != false {
		t.Error("DebugMode should be false")
	}

	hello := 1

	ExecuteInDebugMode(func() {
		hello = 2
	})

	if hello != 1 {
		t.Error("hello should be 1 because DebugMode is false")
	}

	SetDebugMode(true)
	debug = GetDebugMode()

	if debug != true {
		t.Error("DebugMode should be true")
	}

	hello = 1

	ExecuteInDebugMode(func() {
		hello = 2
	})

	if hello != 2 {
		t.Error("hello should be 2 because DebugMode is true")
	}
}

func TestExecuteAndPassthrough(t *testing.T) {
	test := func(args []any) []any {
		t := 0
		for _, v := range args {
			t += v.(int)
		}
		return []any{t}
	}
	testdata := []any{1, 2, 3}
	result := ExecuteAndPassthrough(test, testdata...)
	if len(result) != 1 {
		t.Error("result should be 1 element")
	}
	if result[0] != 6 {
		t.Error("result should be 6")
	}
}
