package interfaces

type IAdd[T any] interface {
	Add(T) T
}

type IAddInplace[T any] interface {
	Add(T)
}

type IAddAny interface {
	Add(any) any
}

type IAddAnyInplace interface {
	Add(any)
}

type IAddAnyFunc interface {
	AddAny(any) any
}

type IAddAnyFuncInplace interface {
	AddAny(any)
}

type ISubstract[T any] interface {
	Substract(T) T
}

type ISubstractInplace[T any] interface {
    Substract(T)
}

type ISubstractAny interface {
	Substract(any) any
}

type ISubstractAnyInplace interface {
	Substract(any)
}

type ISubstractAnyFunc interface {
	SubstractAny(any) any
}

type ISubstractAnyFuncInplace interface {
	SubstractAny(any)
}

type IMultiply[T any] interface {
    Multiply(T) T
}

type IMultiplyInplace[T any] interface {
    Multiply(T)
}

type IMultiplyAny interface {
	Multiply(any) any
}

type IMultiplyAnyInplace interface {
	Multiply(any)
}

type IMultiplyAnyFunc interface {
	MultiplyAny(any) any
}

type IMultiplyAnyFuncInplace interface {
	MultiplyAny(any)
}

type IDivide[T any] interface {
    Divide(T) T
}

type IDivideInplace[T any] interface {
    Divide(T)
}

type IDivideAny interface {
	Divide(any) any
}

type IDivideAnyInplace interface {
	Divide(any)
}

type IDivideAnyFunc interface {
	DivideAny(any) any
}

type IDivideAnyFuncInplace interface {
	DivideAny(any)
}

type IMod[T any] interface {
    Mod(T) T
}

type IModInplace[T any] interface {
    Mod(T)
}

type IModAny interface {
	Mod(any) any
}

type IModAnyInplace interface {
	Mod(any)
}

type IModAnyFunc interface {
	ModAny(any) any
}

type IModAnyFuncInplace interface {
	ModAny(any)
}
