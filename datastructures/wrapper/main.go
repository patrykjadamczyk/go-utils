package wrapper

import "sync"

type Wrapper[T any] struct {
	Value T
}

type MutexedWrapper[T any] struct {
	*Wrapper[T]
	Mutex sync.Mutex
}

type RWMutexedWrapper[T any] struct {
	*Wrapper[T]
	Mutex sync.RWMutex
}
