package base

// Convert basic go function return value to result
func ConvertToResult[T any](value T, err error) Result[T] {
	return Result[T]{DataValue: value, ErrorValue: err}
}

// Convert multiple return value function return in go style to result
func ConvertToResultMultiple(args ...any) Result[any] {
	if len(args) == 0 {
		return MakeOkResult[any](nil)
	}
	if len(args) == 1 {
		if err, ok := args[0].(error); ok {
			if IsError(err) {
				return MakeErrorResult[any](err)
			}
		}
		return MakeOkResult[any](args[0])
	}
	if err2, ok2 := args[len(args)-1].(error); ok2 {
		if IsError(err2) {
			return MakeErrorResultWithValue[any](args, err2)
		}
	}
	return MakeOkResult[any](args)
}

// Check All Results and return first error or last ok value
func CheckAllResults[T any](results ...Result[T]) Result[T] {
	for _, result := range results {
		if result.IsError() {
			return result
		}
	}
	return results[len(results)-1]
}
