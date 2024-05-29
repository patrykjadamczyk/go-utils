package base

type AnyType struct {
	Value any
}

func Any(v any) AnyType {
	return AnyType{Value: v}
}

func (a AnyType) IsNil() bool {
	return a.Value == nil
}

func (a AnyType) IsNotNil() bool {
	return !a.IsNil()
}

func (a AnyType) AsAny() any {
	return a.Value
}

func (a AnyType) IsNull() bool {
	if a.IsNil() {
		return true
	}
	an := UniversalNull(a.Value)
	if an.IsError() {
		return true
	}
	return a.Value == an.GetValue()
}

func (a AnyType) IsNotNull() bool {
	return !a.IsNull()
}

func (a AnyType) Get() any {
	return a.Value
}
