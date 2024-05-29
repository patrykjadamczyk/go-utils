package interfaces

type IToArray[T any] interface {
	ToArray() []T
}

type IToArrayAny interface {
    ToArray() []any
}

type IToArrayAnyFunc interface {
    ToArrayAny() []any
}
