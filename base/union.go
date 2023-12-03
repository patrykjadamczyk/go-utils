package base

import (
	errcat "github.com/patrykjadamczyk/go-utils/exceptions/categories"
)

type UnionInterface interface {
	Get() (uint8, any)
	Set(uint8, any)

	// GetTX() TX
	// SetTX(TX)
}

type UnionTypeConstraint[T1 any, T2 any] interface {
	*T1 | *T2
}

type Union[T1 any, T2 any] struct {
	Data Tuple2[T1, T2]
	Type uint8
}

func (u Union[T1, T2]) GuardInit() {
	AssertThatTyped2[UnionInterface, Union[T1, T2]]().ImplementsT1(u)
	AssertThatTyped2[GuardedObjectInterface, Union[T1, T2]]().ImplementsT1(u)
}

func (u Union[T1, T2]) Get() (uint8, any) {
	switch u.Type {
	case 1:
		return u.Type, any(u.Data.V1)
	case 2:
		return u.Type, any(u.Data.V2)
	default:
		return u.Type, any(nil)
	}
}

func (u Union[T1, T2]) GetT1() T1 {
	return u.Data.V1
}

func (u Union[T1, T2]) GetT2() T2 {
	return u.Data.V2
}

func (u *Union[T1, T2]) SetT1(v T1) {
	u.Type = 1
	u.Data.V1 = v
}

func (u *Union[T1, T2]) SetT2(v T2) {
	u.Type = 2
	u.Data.V2 = v
}

func (u *Union[T1, T2]) Set(t uint8, v any) {
	switch t {
	case 1:
		u.Type = t
		u.Data.V1 = EnsureType[T1](v)
	case 2:
		u.Type = t
		u.Data.V2 = EnsureType[T2](v)
	default:
		err := MakeException(errcat.TypeError, "Union", "Type not supported. Check parameter t if has correct value.")
		panic(err)
	}
}
