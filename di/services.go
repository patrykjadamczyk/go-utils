package di

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

type DiServiceInterface interface {
	Healthcheck() Result[string]
}

type DiServiceContainer struct {
	services map[string]DiServiceInterface
}

func (sc *DiServiceContainer) SetService(name string, service DiServiceInterface) {
	sc.services[name] = service
}

func (sc *DiServiceContainer) GetService(name string) Nullable[DiServiceInterface] {
	if service, ok := sc.services[name]; ok {
		return Some(service)
	}
	return Null[DiServiceInterface]()
}

func (sc *DiServiceContainer) HealthcheckAllServices() bool {
	for _, service := range sc.services {
		if service.Healthcheck().IsError() {
			return false
		}
	}

	return true
}
