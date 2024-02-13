package base

// Pack values provided as arguments to array
func PackToArray(values ...any) []any {
	return values
}

// Explode array into provided references to variables
func Explode(array []any, values ...*any) {
	for i, v := range array {
		*values[i] = v
	}
}

// Explode array into provided references to variables with specified type T
func ExplodeTyped[T any](array []T, values ...*T) {
	for i, v := range array {
		*values[i] = v
	}
}

// Convert specified array of type T to array of type any
func ConvertToAnyArray[T any](array []T) []any {
	result := make([]any, 0)
	for _, v := range array {
		result = append(result, v)
	}
	return result
}
