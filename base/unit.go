package base

import "fmt"

// Note that some value has specific unit like seconds
type UnitValue[T any] struct {
	Value T
	UnitName string
}

// Get Unit Name
func (u UnitValue[T]) GetUnitName() string {
	return u.UnitName
}

// Get Value
func (u UnitValue[T]) GetValue() T {
	return u.Value
}

// String Convert value to string
func (u UnitValue[T]) String() string {
	return fmt.Sprintf("%s", any(u.Value))
}

// Get Go String
func (u UnitValue[T]) GoString() string {
	var ref T
	return fmt.Sprintf(
		"base.Nullable[%T]{%#v %s}",
		ref,
		u.Value,
		u.UnitName,
	)
}

func Unit[T any](value T, unitName string) UnitValue[T] {
	return UnitValue[T]{Value: value, UnitName: unitName}
}
