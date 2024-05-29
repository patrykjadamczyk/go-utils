package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) Add(v T) {
	w.Set(EnsureType[T](add(w.Value, v)))
}

func (w *Wrapper[T]) AddAny(v any) {
	w.Set(EnsureType[T](add(w.Value, v)))
}

func (w *Wrapper[T]) Substract(v T) {
	w.Set(EnsureType[T](substract(w.Value, v)))
}

func (w *Wrapper[T]) SubstractAny(v any) {
	w.Set(EnsureType[T](substract(w.Value, v)))
}

func (w *Wrapper[T]) Multiply(v T) {
	w.Set(EnsureType[T](multiply(w.Value, v)))
}

func (w *Wrapper[T]) MultiplyAny(v any) {
	w.Set(EnsureType[T](multiply(w.Value, v)))
}

func (w *Wrapper[T]) Divide(v T) {
	w.Set(EnsureType[T](divide(w.Value, v)))
}

func (w *Wrapper[T]) DivideAny(v any) {
	w.Set(EnsureType[T](divide(w.Value, v)))
}

func (w *Wrapper[T]) Mod(v T) {
	w.Set(EnsureType[T](mod(w.Value, v)))
}

func (w *Wrapper[T]) ModAny(v any) {
	w.Set(EnsureType[T](mod(w.Value, v)))
}
