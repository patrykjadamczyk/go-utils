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
		"Unit[%T]{%#v %s}",
		ref,
		u.Value,
		u.UnitName,
	)
}

// Note that some value is in specific unit like seconds
func Unit[T any](value T, unitName string) UnitValue[T] {
	return UnitValue[T]{Value: value, UnitName: unitName}
}

// Make function that will return unit
// This is factory of method to make function for specific unit
func MakeUnit[T any](unitName string) func(T) UnitValue[T] {
	return func(value T) UnitValue[T] {
		return Unit[T](value, unitName)
	}
}

// Make function that will normalize to unit type
func NormalizeUnit[T any](unitName string) func(any) Result[UnitValue[T]] {
	return func(value any) Result[UnitValue[T]] {
		if uv, ok := value.(UnitValue[T]); ok {
			if uv.GetUnitName() == unitName {
				return Ok[UnitValue[T]](uv)
			}
			return Err[UnitValue[T]](NewError("Invalid Unit"))
		}
		if v, ok := value.(T); ok {
			return Ok[UnitValue[T]](Unit[T](v, unitName))
		}
		return Err[UnitValue[T]](NewError("Invalid Type"))
	}
}
