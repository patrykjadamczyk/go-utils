package interfaces

type IGetAndSet[T any] interface {
	Get() T
	Set(T)
}

type IGetAndSetAny interface {
	Get() any
	Set(any)
}

type IGetAndSetAnyFunc interface {
	GetAny() any
	SetAny(any)
}
