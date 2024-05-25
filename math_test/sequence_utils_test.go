package math_test

import (
	"testing"

	. "github.com/patrykjadamczyk/go-utils/math"
)

func TestSequenceUtils(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Error("int Sum(2, 3) != 5")
	}
	if Product(2, 3) != 6 {
        t.Error("int Product(2, 3) != 6")
    }
}
