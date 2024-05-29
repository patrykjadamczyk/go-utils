package wrapper

func (w *Wrapper[T]) Get() T {
	return w.Value
}

func (w *Wrapper[T]) GetAny() any {
	return w.Value
}

func (w *Wrapper[T]) GetValue() T {
	return w.Value
}
