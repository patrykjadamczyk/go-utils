package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestPromise(t *testing.T) {
	p := MakePromise(func(v int) int {
		return v + 1
	})
	pr := p.Call(1)
	p.Wait()
	pd := pr.Run()
	if len(pd) != 2 {
		t.Error("Promise result should have 2 element")
	}
	if EnsureType[[]any](pd[0])[0] != 2 {
		t.Error("Promise result should be 2")
	}
	if pd[1] != nil {
		t.Error("Promise result should be nil")
	}
}

func TestPromiseAll(t *testing.T) {
	p1 := func() int {
		return 2
	}
	results, errors := PromiseAll(p1, p1)
	if len(errors) != len(results) {
		t.Error("PromiseAll should have error values and result values equal in length")
	}
	if len(errors) != 2 {
		t.Error("PromiseAll should have 2 elements")
	}
	for _, err := range errors {
		if err != nil {
			t.Error("PromiseAll should not have errors", err)
		}
	}
	for _, val := range results {
		v := EnsureType[[]any](val)
		if len(v) != 1 {
			t.Error("PromiseAll should contain only value 2 so it's value should have only one element", val)
		}
		if v[0] != 2 {
			t.Error("PromiseAll should contain only value 2", val)
		}
	}
}
