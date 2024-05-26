package interfaces

type IGet[T any] interface {
	Get() T
}

type IAdd[T any] interface {
	Add(T) T
}

type IAddInplace[T any] interface {
	Add(T)
}

type ISubstract[T any] interface {
	Substract(T) T
}

type ISubstractInplace[T any] interface {
    Substract(T)
}

type IMultiply[T any] interface {
    Multiply(T) T
}

type IMultiplyInplace[T any] interface {
    Multiply(T)
}

type IDivide[T any] interface {
    Divide(T) T
}

type IDivideInplace[T any] interface {
    Divide(T)
}

type IMod[T any] interface {
    Mod(T) T
}

type IModInplace[T any] interface {
    Mod(T)
}

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

type ILessThanOrEqual[T any] interface {
    LessThanOrEqual(T) bool
}

type IGreaterThan[T any] interface {
    GreaterThan(T) bool
}

type IGreaterThanOrEqual[T any] interface {
    GreaterThanOrEqual(T) bool
}

type ISet[T any] interface {
    Set(T)
}

type IGetAndSet[T any] interface {
    Get() T
    Set(T)
}

type IToString interface {
	ToString() string
}

type IString interface {
	String() string
}

type ILen[T any] interface {
    Len() T
}

type IToTuple2[T1, T2 any] interface {
    ToTuple() (T1, T2)
}

type IToTuple3[T1, T2, T3 any] interface {
    ToTuple() (T1, T2, T3)
}

type IToTuple4[T1, T2, T3, T4 any] interface {
    ToTuple() (T1, T2, T3, T4)
}

type IFromString interface {
	FromString(string)
}

type IFromStringToObj[T any] interface {
	FromString(string) T
}

type IContains[T any] interface {
	Contains(T) bool
}

type IInsert[T any] interface {
    Insert(T) T
}

type IAppend[T any] interface {
    Append(T) T
}

type IInsertInplace[T any] interface {
    Insert(T)
}

type IAppendInplace[T any] interface {
    Append(T)
}
