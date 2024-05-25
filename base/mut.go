package base

// MutableRefValue is structure for Mutable Value that acts like reference
type MutableRefValue[T any] struct {
	Data *T
}

// Set value
func (mrv *MutableRefValue[T]) Set(value T) {
	if mrv.Data == nil {
		mrv.Data = &value
		return
	}
	*mrv.Data = value
}

// Get value
func (mrv MutableRefValue[T]) Get() T {
	if mrv.Data == nil {
		return Null[T]().Data
	}
	return *mrv.Data
}

// Get pointer to value
func (mrv *MutableRefValue[T]) GetPtr() *T {
	return mrv.Data
}

// MakeRefValue makes mutable reference to specified value still acting completely like value
func MutRefValue[T any](value T) MutableRefValue[T] {
	// data := make([]T, 0)
	// data = append(data, value)
	return MutableRefValue[T]{Data: &value}
}
