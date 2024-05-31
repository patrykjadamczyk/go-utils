package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	compare "github.com/patrykjadamczyk/go-utils/base/CompareEnum"
)

func (w *Wrapper[T]) Equals(v T) bool {
	return UniversalEquals(w.Get(), v).GetValue()
}

func (w *Wrapper[T]) EqualsAny(v any) bool {
	return UniversalEquals(w.Get(), v).GetValue()
}

func (w *Wrapper[T]) NotEquals(v T) bool {
	return !w.Equals(v)
}

func (w *Wrapper[T]) NotEqualsAny(v any) bool {
	return !w.EqualsAny(v)
}

func (w *Wrapper[T]) LessThan(v T) bool {
	return UniversalCompare(w.Get(), v) == compare.Less
}

func (w *Wrapper[T]) LessThanAny(v any) bool {
	return UniversalCompare(w.Get(), v) == compare.Less
}

func (w *Wrapper[T]) LessThanOrEqual(v T) bool {
	c := UniversalCompare(w.Get(), v)
	return c == compare.Less || c == compare.Equal
}

func (w *Wrapper[T]) LessThanOrEqualAny(v any) bool {
	c := UniversalCompare(w.Get(), v)
	return c == compare.Less || c == compare.Equal
}

func (w *Wrapper[T]) GreaterThan(v T) bool {
	return UniversalCompare(w.Get(), v) == compare.Greater
}

func (w *Wrapper[T]) GreaterThanAny(v any) bool {
	return UniversalCompare(w.Get(), v) == compare.Greater
}

func (w *Wrapper[T]) GreaterThanOrEqual(v T) bool {
	c := UniversalCompare(w.Get(), v)
	return c == compare.Greater || c == compare.Equal
}

func (w *Wrapper[T]) GreaterThanOrEqualAny(v any) bool {
	c := UniversalCompare(w.Get(), v)
	return c == compare.Greater || c == compare.Equal
}
