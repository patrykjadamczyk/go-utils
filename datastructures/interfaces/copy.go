package interfaces

type ICopy[T any] interface {
	Copy() T
}

type ICopyAny interface {
	Copy() any
}

type ICopyAnyFunc interface {
    CopyAny() any
}

type IClone[T any] interface {
	Clone() T
}

type ICloneAny interface {
	Clone() any
}

type ICloneAnyFunc interface {
    CloneAny() any
}
