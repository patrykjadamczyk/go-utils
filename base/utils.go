package base

func Range(end int) []int {
	return make([]int, end)
}

func StartEndRange(start int, end int) []int {
	return ComplexRange(start, end, 1)
}

func StartEndNonInclusiveRange(start int, end int) []int {
	return ComplexNonInclusiveRange(start, end, 1)
}

func ComplexRange(start int, end int, step int) []int {
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

func ComplexNonInclusiveRange(start int, end int, step int) []int {
	var result []int
	for i := start; i < end; i += step {
		result = append(result, i)
	}
	return result
}
