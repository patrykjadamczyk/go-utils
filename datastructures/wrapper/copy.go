package wrapper

func (w *Wrapper[T]) Copy() Wrapper[T] {
    return Wrapper[T]{ Value: w.Value }
}

func (w *Wrapper[T]) CopyAny() any {
    return Wrapper[T]{ Value: w.Value }
}

func (w *Wrapper[T]) Clone() Wrapper[T] {
    return Wrapper[T]{ Value: w.Value }
}

func (w *Wrapper[T]) CloneAny() any {
    return Wrapper[T]{ Value: w.Value }
}
