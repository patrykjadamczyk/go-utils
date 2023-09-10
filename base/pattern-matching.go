package base

type SwitchResult[SV any, V any] struct {
	SwitchValue SV
	Value       V
	SwitchFound bool
}

func Switch[SV any, V any](value SV) *SwitchResult[SV, V] {
	return &SwitchResult[SV, V]{
		SwitchValue: value,
		SwitchFound: false,
	}
}

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

func (s *SwitchResult[SV, V]) Default(f func(SV) V) *SwitchResult[SV, V] {
	if s.SwitchFound {
		return s
	}
	s.Value = f(s.SwitchValue)
	s.SwitchFound = true
	return s
}

func Ternary[V any](condition bool, ifTrue V, ifFalse V) V {
	if condition {
		return ifTrue
	}
	return ifFalse
}
