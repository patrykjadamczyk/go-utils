package base

// Switch Result Type
type SwitchResult[SV any, V any] struct {
	// Value to Switch On
	SwitchValue SV
	// Result Value
	Value       V
	// Is Switch Found
	SwitchFound bool
}

// Switch Statement based on provided value
func Switch[SV any, V any](value SV) *SwitchResult[SV, V] {
	return &SwitchResult[SV, V]{
		SwitchValue: value,
		SwitchFound: false,
	}
}

// Case for Switch
// specified function have to take switch value and return tuple of value and bool if it was found in that case
func (s *SwitchResult[SV, V]) Case(f func(SV) (V, bool)) *SwitchResult[SV, V] {
	if s.SwitchFound {
		return s
	}
	if v, ok := f(s.SwitchValue); ok {
		s.Value = v
		s.SwitchFound = true
		return s
	}
	return s
}

// Default Case for Switch
// specified function have to take switch value and return value
func (s *SwitchResult[SV, V]) Default(f func(SV) V) *SwitchResult[SV, V] {
	if s.SwitchFound {
		return s
	}
	s.Value = f(s.SwitchValue)
	s.SwitchFound = true
	return s
}

// Ternary Operation
// You can see this function as something like condition ? ifTrue : ifFalse
func Ternary[V any](condition bool, ifTrue V, ifFalse V) V {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// Ternary that checks if value provided is nil if it's nil it uses second argument as return
// You can see this function as something like value ?? ifNullValue
func TernaryNull[V any](value V, ifNullValue V) V {
	if any(value) == nil {
		return ifNullValue
	}
	return value
}

// Ternary that checks if value provided is not null according to NullableType if it's null it uses second argument as
// return
// You can see this function as something like value ?? ifNullValue but a bit more complicated because it's checks Value
// for OptionType (Nullable) if it is something or nothing
func TernaryNullable[V any](value Nullable[V], ifNullValue V) V {
	if value.IsError() {
		return ifNullValue
	}
	return value.Data
}
