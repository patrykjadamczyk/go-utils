package base

// Slice type
type Slice[T any] []T

// Len returns the number of elements in Slice, if Slice is nil then Len is 0.
func (s Slice[T]) Len() int {
    return len(s)
}

// Cap returns the capacity of Slice, if Slice is nil then Cap is 0.
func (s Slice[T]) Cap() int {
    return cap(s)
}

// Append appends ts to the end of s and returns the updated slice.
func (s Slice[T]) Append(ts ...T) Slice[T] {
    return append(s, ts...)
}

// Delete removes the element at index i from s and returns the updated slice.
func (s Slice[T]) Delete(i int) Slice[T] {
	sc := s.Copy()
    return append(sc[:i], sc[i+1:]...)
}

// Make copy of slice
func (s Slice[T]) Copy() Slice[T] {
	var ret = make(Slice[T], 0)
	ret = append(ret, s...)
	return ret
}
