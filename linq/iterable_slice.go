package linq

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

type IterableSlice[TV any] struct {
	slice []TV
	index int
}

func (is *IterableSlice[TV]) Next() (TV, bool) {
	is.index++
	if is.index >= len(is.slice) {
		return None[TV]().ValueOrZero(), false
	}
	return is.slice[is.index-1], true
}
