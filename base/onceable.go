package base

// Onceable Object
type Onceable[T any] struct {
	value       T
	onceHandler func() T
	isSet       bool
}

// Get value of Onceable
func (o *Onceable[T]) Get() T {
	if !o.isSet {
		if o.onceHandler == nil {
			panic("Onceable Error: OnceHandler is not set")
		}
		o.value = o.onceHandler()
		o.isSet = true
	}
	return o.value
}

// Get value of Onceable as any
func (o Onceable[T]) GetAny() any {
	return o.Get()
}

// Make New Onceable
func MakeOnceable[T any](handler func() T) Onceable[T] {
	return Onceable[T]{onceHandler: handler}
}
