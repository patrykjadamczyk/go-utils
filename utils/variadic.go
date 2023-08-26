package utils

func MultipleValuesToArray(values ...any) []any {
    return values
}

func ArrayToMultipleValues(array []any, values ...*any) {
    for i, v := range array {
        *values[i] = v
    }
}

func ArrayToMultipleValuesTyped[T any](array []T, values ...*T) {
    for i, v := range array {
        *values[i] = v
    }
}

func AnyArray[T any](array []T) []any {
    result := make([]any, len(array))
    for _, v := range array {
        result = append(result, v)
    }
    return result
}
