package base

// MutableRefValue is structure for Mutable Value that acts like reference
type MutableRefValue[T any] struct {
	Data []T
}

// Set value
func (mrv *MutableRefValue[T]) Set(value T) {
	if mrv.Data == nil || len(mrv.Data) == 0 {
		mrv.Data = make([]T, 0)
		mrv.Data = append(mrv.Data, value)
		return
	}
	mrv.Data[0] = value
}

// Get value
func (mrv MutableRefValue[T]) Get() T {
	if mrv.Data == nil || len(mrv.Data) == 0 {
		return Null[T]().Data
	}
	return mrv.Data[0]
}

// Get pointer to value
func (mrv *MutableRefValue[T]) GetPtr() *T {
	if mrv.Data == nil || len(mrv.Data) == 0 {
		return nil
	}
	return &mrv.Data[0]
}

// MakeRefValue makes mutable reference to specified value still acting completely like value
func MutRefValue[T any](value T) MutableRefValue[T] {
	data := make([]T, 0)
	data = append(data, value)
	return MutableRefValue[T]{Data: data}
}
