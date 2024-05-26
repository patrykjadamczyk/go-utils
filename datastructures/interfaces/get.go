package interfaces

type IGet[T any] interface {
	Get() T
}

type IGetAny interface {
	Get() any
}

type IGetAnyFunc interface {
	GetAny() any
}

type IGetValue[T any] interface {
	GetValue() T
}

type IGetValueAny interface {
    GetValue() any
}
