package bool

// Returns true if both a and b are true
func And(a, b bool) bool {
	return a && b
}

// Returns true if either a or b is true
func Or(a, b bool) bool {
	return a || b
}

// Returns true if a and b are different
func Xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

// Returns true if a and b are the same
func Xnor(a, b bool) bool {
	return (a && b) || (!a && !b)
}

// Returns the opposite of a
func Not(a bool) bool {
	return !a
}

// Returns true if either a or b is false
func Nand(a, b bool) bool {
	return !a || !b
}

// Returns true if both a and b are false
func Nor(a, b bool) bool {
	return !a && !b
}

// Returns 1 if a is true, 0 if a is false
func ToInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

// Returns "True" if a is true, "False" if a is false
func ToString(a bool) string {
	if a {
		return "True"
	}
	return "False"
}

// Run a callback function if the given bool is false, otherwise return the given default value
func Guard[T any](when bool, default_value T, alternative func() T) T {
	if when {
		return default_value
	}
	return alternative()
}

// Run a callback function if the given bool is true, otherwise return the given alternative callback
func LazyGuard[T any](when bool, default_value func() T, alternative func() T) T {
	if when {
		return default_value()
	}
	return alternative()
}
