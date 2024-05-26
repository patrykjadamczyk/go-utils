package base

// ContextManagerInterface with methods on entering and exiting context
// This interface could be on pointer type or struct itself based if it have to manipulate data on struct
type ContextManagerInterface interface {
	// Enter into Context
	Enter()
	// Exit from Context
	Exit()
}

// GoContextManagerInterface with methods on entering and exiting context
// This interface is based on close method that is made in some go objects for use in defer
// This interface could be on pointer type or struct itself based if it have to manipulate data on struct
type GoContextManagerInterface interface {
	// Close on Context Leave
	Close()
}

// Do something with context from context manager object
// Before executing function specified it calls Enter and Exit is called on defer
func WithContext[T any](
	ctx ContextManagerInterface,
	f func(ctx ContextManagerInterface) T,
) T {
	defer ctx.Exit()
	ctx.Enter()
	return f(ctx)
}

// Do something with context from context manager object
// Before executing function specified it calls Close on defer
func WithGoContext[T any](
	ctx GoContextManagerInterface,
	f func(ctx GoContextManagerInterface) T,
) T {
	defer ctx.Close()
	return f(ctx)
}
