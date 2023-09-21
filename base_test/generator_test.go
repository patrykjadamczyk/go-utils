package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestGenerator(t *testing.T) {
	state := 0
	stopGeneratorChan := make(chan struct{})
	genResults := MakeGenerator[int](stopGeneratorChan, func(ch chan int) bool {
		state++
		Yield(ch, state)
		if state == 3 {
			return false
		}
		return true
	})
	state2 := 0
	for v := range genResults {
		state2++
		if v != state2 {
			t.Errorf("state[%d] should be same as state2[%d]", state, state2)
		}
	}
}
