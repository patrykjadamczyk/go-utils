package base

// Unwrappable Interface Object
type UnwrappableInterface[V any] interface {
	// Get underlying value or panic if error
	Unwrap() V
	// Get underlying value or return default value if error is found
	UnwrapOr(defaultVal V) V
	// Unwrap but if error occurs run specified function f instead of panic
	// Return underlying value
	UnwrapOrErr(f func(UnwrappableInterface[V])) V
	// Unwraps as result type for if error checks and run specified function f
	// instead of panic
	// Return result of underlying value (having err if it's found)
	UnwrapAsResultOrErr(f func(UnwrappableInterface[V])) Result[V]
	// Unwrap value and error separately (Result -> Go normal returns)
	UnwrapWithErr() (V, error)
	// Expect correct value if error is found panic with specified message
	Expect(err any)
}

// Errorable Generic Result Interface Object
// This interface is meant to be used for checking if Result[T] has error or similiar object to Result[T] without option
// to check for underlying type of Result
type ErrorableGenericResultInterface interface {
	// Check if Object has error
	IsError() bool
	// Get Error if Object has error or nil if not
	GetError() error
}
