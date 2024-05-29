package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *MutexedWrapper[T]) Set(v T) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = v
}

func (w *MutexedWrapper[T]) SetAny(v any) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = EnsureType[T](v)
}

func (w *MutexedWrapper[T]) SetValue(v T) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = v
}

func (w *RWMutexedWrapper[T]) Set(v T) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = v
}

func (w *RWMutexedWrapper[T]) SetAny(v any) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = EnsureType[T](v)
}

func (w *RWMutexedWrapper[T]) SetValue(v T) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Value = v
}

func (w *RWMutexedWrapper[T]) Get() T {
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()
	return w.Value
}

func (w *RWMutexedWrapper[T]) GetAny() any {
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()
	return w.Value
}

func (w *RWMutexedWrapper[T]) GetValue() T {
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()
	return w.Value
}
