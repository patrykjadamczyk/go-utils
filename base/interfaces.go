package base

// Unwrappable Interface Object
type UnwrappableInterface[V any] interface {
	// Get underlying value or panic if error
	Unwrap() V
	// Get underlying error or panic if correct value
	UnwrapErr() error
	// Get underlying value or return default value if error is found
	UnwrapOr(defaultVal V) V
	// Unwrap value and error separately (Result -> Go normal returns)
	UnwrapWithErr() (V, error)
	// Expect correct value if error is found panic with specified message
	Expect(err any)
	// Expect error value if error is not found panic with specified message
	ExpectErr(err any)
}

// Unwrappable Interface Extension
// This functions should be implemented by any object that want to implement Unwrappable Interface
// This functions are based on UnwrappableInterface itself being a problem for type checks for some reason
type ExtendedUnwrappableInterface[V any] interface {
	// Unwrap but if error occurs run specified function f instead of panic
	// Return underlying value
	UnwrapOrErr(f func(UnwrappableInterface[V])) V
	// Unwraps as result type for if error checks and run specified function f
	// instead of panic
	// Return result of underlying value (having err if it's found)
	UnwrapAsResultOrErr(f func(UnwrappableInterface[V])) Result[V]
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

// Defaultable Interface
type DefaultableInterface[T any] interface {
	// Get Default Value
	GetDefault() T
}

// Assertable Interface Object
type AssertableInterface interface {
    // Method that checks if object is valid
    // This method should use Assert method inside or AssertCustomError method
    Assert()
}

type AssertableBoolInterface interface {
    // Method that checks if object is valid
    // This method should return bool if object is valid
    AssertBool() bool
}

// Guarded Object Interface
// Guarded Object is meant to be object that has some guards that need to be the case or this object is not valid at all
type GuardedObjectInterface interface {
    // Initialize method for Guarded Object
    // This method is meant to be used in init method of module and contain only assert that check for example if object
    // is compatible with some interfaces or some version of go or something else and things like that for objects logic
    // to be correct
    GuardInit()
}
