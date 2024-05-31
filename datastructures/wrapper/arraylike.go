package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) Len() int {
	return UniversalLen(w.Get())
}

func (w *Wrapper[T]) Contains(v T) bool {
	wa := w.ToArray()
	for _, wv := range wa {
		if UniversalEquals(wv, v).GetValue() {
			return true
		}
	}
	return false
}

func (w *Wrapper[T]) ContainsAny(v any) bool {
	wa := w.ToArray()
	for _, wv := range wa {
		if UniversalEquals(wv, v).GetValue() {
			return true
		}
	}
	return false
}

func makeArrayWrapper[T any](v []T) Wrapper[[]any] {
	res := make([]any, 0)
	for _, item := range EnsureType[[]T](v) {
		res = append(res, item)
	}
	return Wrapper[[]any]{Value: res}
}

// Result of that function you need to use through EnsureType[Wrapper[[]T]]()
func (w *Wrapper[T]) Insert(v T) any {
	var ret = append(w.ToArray(), v)
	return makeArrayWrapper(ret)
}

func (w *Wrapper[T]) InsertAny(v any) any {
	var ret = append(w.ToArray(), EnsureType[T](v))
	return makeArrayWrapper(ret)
}

// Result of that function you need to use through EnsureType[Wrapper[[]T]]()
func (w *Wrapper[T]) Append(v T) any {
	var ret = append(w.ToArray(), v)
	return makeArrayWrapper(ret)
}

func (w *Wrapper[T]) AppendAny(v any) any {
	var ret = append(w.ToArray(), EnsureType[T](v))
	return makeArrayWrapper(ret)
}

// Result of that function you need to use through EnsureType[Wrapper[[]T]]()
func (w *Wrapper[T]) Prepend(v T) any {
	var ret = make([]T, 0)
	ret = append(ret, v)
	ret = append(ret, w.ToArray()...)
	return makeArrayWrapper(ret)
}

func (w *Wrapper[T]) PrependAny(v any) any {
	var ret = make([]T, 0)
	ret = append(ret, EnsureType[T](v))
	ret = append(ret, w.ToArray()...)
	return makeArrayWrapper(ret)
}
