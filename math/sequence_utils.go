package math

func Sum[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint](values ...T) T {
	var sum T = 0
	for _, value := range values {
		sum += value
    }
    return sum
}

func Product[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint](values ...T) T {
	var product T = 0
	for _, value := range values {
		product *= value
    }
    return product
}
