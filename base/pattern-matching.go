package base

// Switch Result Type
type SwitchResult[SV any, V any] struct {
	// Value to Switch On
	SwitchValue SV
	// Result Value
	Value V
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

// Ternary that checks if value provided is nil if it's nil it uses second argument as return
// You can see this function as something like value ?? ifNullValue
// This function checks for nil using IsNil function
func TernaryNil[V any](value V, ifNullValue V) V {
	if IsNil(value) {
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

// Ternary that checks if result is ok or err and returns specified value for this condition
func TernaryResult[V any](result Result[V], ifOk V, ifErr V) V {
	if result.IsError() {
		return ifErr
	}
	return ifOk
}

// Return the first valid value or null zero value if not found
func Coalesce[V any](values ...V) V {
	for _, v := range values {
		switch any(v).(type) {
		case ErrorableGenericResultInterface:
			if !EnsureType[ErrorableGenericResultInterface](v).IsError() {
				return v
			}
		default:
			if !IsNil(v) {
				return v
			}
		}
	}

	return Null[V]().ValueOrZero()
}

// Map Switch Result Type
type MapSwitchResult[SV comparable, V any] struct {
	ValueMap   map[SV]func() V
	DefaultMap func() V
}

// Make Map Switch Statement
func MapSwitch[SV comparable, V any]() *MapSwitchResult[SV, V] {
	return &MapSwitchResult[SV, V]{
		ValueMap: make(map[SV]func() V),
	}
}

// Make case for Map switch
// specified function will be called only if switch value will be same as specified value
func (s *MapSwitchResult[SV, V]) Case(switchValue SV, f func() V) *MapSwitchResult[SV, V] {
	if s.ValueMap[switchValue] == nil {
		s.ValueMap[switchValue] = f
	}
	return s
}

// Make default case for Map switch
// specified function will be called only if switch value will not be found
func (s *MapSwitchResult[SV, V]) Default(f func() V) *MapSwitchResult[SV, V] {
	if s.DefaultMap == nil {
		s.DefaultMap = f
	}
	return s
}

// Get value of Map switch
// This will return case for specified value if found, default value if it exists and null value if none cases match
func (s *MapSwitchResult[SV, V]) Get(switchValue SV) V {
	if s.ValueMap[switchValue] == nil {
		if s.DefaultMap == nil {
			return Null[V]().ValueOrZero()
		}
		return s.DefaultMap()
	}
	return s.ValueMap[switchValue]()
}
