package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) Set(v T) {
	w.Value = v
}

func (w *Wrapper[T]) SetAny(v any) {
	w.Value = EnsureType[T](v)
}

func (w *Wrapper[T]) SetValue(v T) {
	w.Value = v
}
