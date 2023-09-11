package base

func Range(end int) []int {
	return make([]int, end)
}

// Range function returns an array to be used with for _, i := range StartEndRange(end)
func StartEndRange(start int, end int) []int {
	return ComplexRange(start, end, 1)
}

// Range function returns an array to be used with for _, i := range StartEndNonInclusiveRange(end)
func StartEndNonInclusiveRange(start int, end int) []int {
	return ComplexNonInclusiveRange(start, end, 1)
}

// Range function returns an array to be used with for _, i := range ComplexRange(end)
func ComplexRange(start int, end int, step int) []int {
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

// Range function returns an array to be used with for _, i := range ComplexNonInclusiveRange(end)
func ComplexNonInclusiveRange(start int, end int, step int) []int {
	var result []int
	for i := start; i < end; i += step {
		result = append(result, i)
	}
	return result
}
