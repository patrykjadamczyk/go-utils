package di

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

type DiContainer struct {
	Services map[string]any
}

func GetService[T any](dc DiContainer, name string) Nullable[T] {
	if dc.Services == nil {
		dc.Services = make(map[string]any)
		return None[T]()
	}

	if service, ok := dc.Services[name]; ok {
		return Some[T](EnsureType[T](service))
	}

	return None[T]()
}

func SetService[T any](dc *DiContainer, name string, service T) {
	if dc.Services == nil {
		dc.Services = make(map[string]any)
	}

	dc.Services[name] = service
}
