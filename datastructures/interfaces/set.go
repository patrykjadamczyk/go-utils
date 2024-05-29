package interfaces

type ISet[T any] interface {
	Set() T
}

type ISetAny interface {
	Set() any
}

type ISetAnyFunc interface {
	SetAny() any
}

type ISetValue[T any] interface {
	SetValue() T
}

type ISetValueAny interface {
	SetValue() any
}
