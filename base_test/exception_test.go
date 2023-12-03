package base_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/tests"
)

func TestException(t *testing.T) {
	t1 := Exception{}
	t2 := ExtendedException{}
	t3 := SubCategoryExtendedException{}

	t1.Init("Hello", "World")
	t2.Init("Hello", "World")
	t3.Init("Hello", "W", "World")

	tests.TestIfNotPanics(t, func() {
		t1.GuardInit()
		t2.GuardInit()
		t3.GuardInit()
	}, "Exceptions should not panic on GuardInit")
}
