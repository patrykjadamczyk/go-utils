package base

// Observable Object
type Observable[T any] struct {
	Value       T
	GetHandlers []func(Observable[T])
	SetHandlers []func(Observable[T])
}

// Set Value of Observable
func (o *Observable[T]) SetValue(value T) {
	o.Value = value
	o.NotifySet()
}

// Get Value of Observable
func (o *Observable[T]) GetValue() T {
	o.NotifyGet()
	return o.Value
}

// Add Get Handler
// Get Handler is called when GetValue is called
func (o *Observable[T]) AddGetHandler(handler func(Observable[T])) {
	o.GetHandlers = append(o.GetHandlers, handler)
}

// Add Set Handler
// Set Handler is called when SetValue is called
func (o *Observable[T]) AddSetHandler(handler func(Observable[T])) {
	o.SetHandlers = append(o.SetHandlers, handler)
}

// Notify Get Handlers
func (o *Observable[T]) NotifyGet() {
	for _, handler := range o.GetHandlers {
		handler(*o)
	}
}

// Notify Set Handlers
func (o *Observable[T]) NotifySet() {
	for _, handler := range o.SetHandlers {
		handler(*o)
	}
}

// Make New Observable
func MakeObservable[T any](value T) Observable[T] {
	return Observable[T]{Value: value}
}
