package base

import "sync"

// Thread Safe Object that builds arround some value and make methods to set it and get it in thread safe way
type ThreadSafeObject[T any] struct {
	obj  T
	lock sync.RWMutex
}

// Set underlying object to new value
func (tso *ThreadSafeObject[T]) Set(newObj T) {
	tso.lock.Lock()
	defer tso.lock.Unlock()
	tso.obj = newObj
}

// Get underlying object value
func (tso *ThreadSafeObject[T]) Get() T {
	tso.lock.RLock()
	defer tso.lock.RUnlock()
	return tso.obj
}

func MakeThreadSafeObject[T any](obj T) ThreadSafeObject[T] {
	return ThreadSafeObject[T]{obj: obj}
}
