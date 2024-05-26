package interfaces

type IToString interface {
	ToString() string
}

type IString interface {
	String() string
}

type IFromString interface {
	FromString(string)
}

type IFromStringToObj[T any] interface {
	FromString(string) T
}

type IFromStringToAny interface {
	FromString(string) any
}
