package stack_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/datastructures/stack"
)

func TestStack(t *testing.T) {
	ts := *NewStack[int]()
	if !ts.IsEmpty() {
		t.Error("New stack should be empty")
	}
	ts.Push(21)
	if ts.Len() != 1 {
		t.Error("Push should add an element")
	}
	if ts.Get()[0] != 21 {
		t.Error("Push should add an element and this element should be 21")
	}
	if ts.Pop() != 21 {
		t.Error("Pop should return the last element and this element should be 21")
	}
	if !ts.IsEmpty() {
		t.Error("Pop should remove last element")
	}
}
