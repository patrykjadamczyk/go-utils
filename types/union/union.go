package union

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/*
You can make correct union type like this
type Union interface {
	T1 | T2
}
*/

type Union[T1 any, T2 any] struct {
	Data any
}

func Value[T1 any, T2 any](value any) Union[T1, T2] {
	switch value.(type) {
	case T1:
		return Union[T1, T2]{Data: value}
	case T2:
		return Union[T1, T2]{Data: value}
	default:
		panic("Invalid Type")
	}
}

func ValueFromPointer[T1 any, T2 any](value *any) Union[T1, T2] {
	return Value[T1, T2](*value)
}

func (u Union[T1, T2]) ExactEqual(other Union[T1, T2]) bool {
	return any(u.Data) == any(other.Data)
}

func (u Union[T1, T2]) Equal(other Union[T1, T2]) bool {
	switch any(u.Data).(type) {
	case time.Time:
		uValue := any(u.Data).(time.Time)
		otherValue := any(other.Data).(time.Time)
		return uValue.Equal(otherValue)
	}
	return u.ExactEqual(other)
}

func (u Union[T1, T2]) String() string {
	return fmt.Sprintf("%s", any(u.Data))
}

func (u Union[T1, T2]) GoString() string {
	var ref1 T1
	var ref2 T2
	return fmt.Sprintf("union.Union[%T,%T]{Data:%#v}", ref1, ref2, u.Data)
}

func (u Union[T1, T2]) Value() (driver.Value, error) {
	return any(u.Data), nil
}
