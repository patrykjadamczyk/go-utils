package interfaces

type IEquals[T any] interface {
    Equals(T) bool
}

type INotEquals[T any] interface {
    NotEquals(T) bool
}

type IEqualsAny interface {
    Equals(any) bool
}

type INotEqualsAny interface {
    NotEquals(any) bool
}

type IEqualsAnyFunc interface {
    EqualsAny(any) bool
}

type INotEqualsAnyFunc interface {
    NotEqualsAny(any) bool
}

type ILessThan[T any] interface {
    LessThan(T) bool
}

type ILessThanAny interface {
    LessThan(any) bool
}

type ILessThanAnyFunc interface {
    LessThanAny(any) bool
}

type ILessThanOrEqual[T any] interface {
    LessThanOrEqual(T) bool
}

type ILessThanOrEqualAny interface {
    LessThanOrEqual(any) bool
}

type ILessThanOrEqualAnyFunc interface {
    LessThanOrEqualAny(any) bool
}

type IGreaterThan[T any] interface {
    GreaterThan(T) bool
}

type IGreaterThanAny interface {
    GreaterThan(any) bool
}

type IGreaterThanAnyFunc interface {
    GreaterThanAny(any) bool
}

type IGreaterThanOrEqual[T any] interface {
    GreaterThanOrEqual(T) bool
}

type IGreaterThanOrEqualAny interface {
    GreaterThanOrEqual(any) bool
}

type IGreaterThanOrEqualAnyFunc interface {
    GreaterThanOrEqualAny(any) bool
}
