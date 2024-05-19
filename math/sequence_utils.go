package math

import "github.com/patrykjadamczyk/go-utils/base"

func Sum[T base.NumericType](values ...T) T {
	var sum T = 0
	for _, value := range values {
		sum += value
    }
    return sum
}

func Product[T base.NumericType](values ...T) T {
	var product T = 0
	for _, value := range values {
		product *= value
    }
    return product
}
