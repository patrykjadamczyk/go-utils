package base

func EnsureType[T any](value any) T {
    switch value.(type) {
    case T:
        return value.(T)
    default:
        panic("Invalid Type")
    }
}
