package base

func PackToArray(values ...any) []any {
    return values
}

func Explode(array []any, values ...*any) {
    for i, v := range array {
        *values[i] = v
    }
}

func ExplodeTyped[T any](array []T, values ...*T) {
    for i, v := range array {
        *values[i] = v
    }
}

func ConvertToAnyArray[T any](array []T) []any {
    result := make([]any, len(array))
    for _, v := range array {
        result = append(result, v)
    }
    return result
}
