package base

// Ensure that value is of type T or panic
func EnsureType[T any](value any) T {
    switch value.(type) {
    case T:
        return value.(T)
    default:
        panic("Invalid Type")
    }
}
