package linq

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"golang.org/x/exp/constraints"
)

func LinqWhere[T any](obj []T, predicate func(T) bool) (result []T) {
	for _, item := range obj {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func LinqWhereIterable[T any, TV IIterable[T]](obj TV, predicate func(T) bool) (result []T) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func LinqSelect[IT any, OT any](obj []IT, predicate func(IT) OT) (result []OT) {
	for _, item := range obj {
		result = append(result, predicate(item))
	}

	return result
}

func LinqSelectIterable[IT any, TV IIterable[IT], OT any](obj TV, predicate func(IT) OT) (result []OT) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		result = append(result, predicate(item))
	}

	return result
}

func LinqTake[T any](obj []T, count int) []T {
	return obj[:count]
}

func LinqTakeIterable[T any, TV IIterable[T]](obj TV, count int) (result []T) {
	for i := 0; i < count; i++ {
		item, ok := obj.Next()
		if !ok {
			break
		}
		result = append(result, item)
	}

	return result
}

func LinqSkip[T any](obj []T, count int) []T {
	return obj[count:]
}

func LinqSkipIterable[T any, TV IIterable[T]](obj TV, count int) (result []T) {
	i := 0

	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if i <= count {
			i++
			continue
		}
		result = append(result, item)
	}

	return result
}

func LinqAll[T any](obj []T, predicate func(T) bool) bool {
	for _, item := range obj {
		if !predicate(item) {
			return false
		}
	}

	return true
}

func LinqAllIterable[T any, TV IIterable[T]](obj TV, predicate func(T) bool) bool {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if !predicate(item) {
			return false
		}
	}

	return true
}

func LinqAny[T any](obj []T, predicate func(T) bool) bool {
	for _, item := range obj {
		if predicate(item) {
			return true
		}
	}

	return false
}

func LinqAnyIterable[T any, TV IIterable[T]](obj TV, predicate func(T) bool) bool {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if predicate(item) {
			return true
		}
	}

	return false
}

func LinqContains[T comparable](obj []T, searchValue T) bool {
	for _, item := range obj {
		if item == searchValue {
			return true
		}
	}

	return false
}

func LinqContainsIterable[T comparable, TV IIterable[T]](obj TV, searchValue T) bool {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if item == searchValue {
			return true
		}
	}

	return false
}

func LinqCount[T any](obj []T) int {
	return len(obj)
}

func LinqCountIterable[T any, TV IIterable[T]](obj TV) int {
	count := 0
	for {
		_, ok := obj.Next()
		if !ok {
			break
		}
		count++
	}

	return count
}

func LinqEach[T any](obj []T, action func(T)) {
	for _, item := range obj {
		action(item)
	}
}

func LinqEachIterable[T any, TV IIterable[T]](obj TV, action func(T)) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		action(item)
	}
}

func LinqFind[T any](obj []T, predicate func(T) bool) Nullable[T] {
	for _, item := range obj {
		if predicate(item) {
			return Some[T](item)
		}
	}

	return None[T]()
}

func LinqFindIterable[T any, TV IIterable[T]](obj TV, predicate func(T) bool) Nullable[T] {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if predicate(item) {
			return Some[T](item)
		}
	}

	return None[T]()
}

func LinqGroupBy[T any, K comparable](obj []T, keyGenerator func(T) K) map[K][]T {
	result := make(map[K][]T, 0)

	for _, item := range obj {
		k := keyGenerator(item)
		if r, ok := result[k]; ok {
			result[k] = append(r, item)
		} else {
			result[k] = []T{item}
		}
	}

	return result
}

func LinqGroupByIterable[T any, TV IIterable[T], K comparable](obj TV, keyGenerator func(T) K) map[K][]T {
	result := make(map[K][]T, 0)

	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		k := keyGenerator(item)
		if r, ok := result[k]; ok {
			result[k] = append(r, item)
		} else {
			result[k] = []T{item}
		}
	}

	return result
}

func LinqToArray[T any](obj []T) []T {
	return obj
}

func LinqToArrayIterable[T any, TV IIterable[T]](obj TV) (result []T) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		result = append(result, item)
	}

	return result
}

func LinqJoin[T any](objs ...[]T) (result []T) {
	for _, obj := range objs {
		result = append(result, obj...)
	}

	return result
}

func LinqJoinIterable[T any, TV IIterable[T]](objs ...TV) (result []T) {
	for _, obj := range objs {
		for {
			item, ok := obj.Next()
			if !ok {
				break
			}
			result = append(result, item)
		}
	}

	return result
}

func LinqLast[T any](obj []T) T {
	return obj[len(obj)-1]
}

func LinqLastIterable[T any, TV IIterable[T]](obj TV) T {
	arr := LinqToArrayIterable[T, TV](obj)
	return arr[len(arr)-1]
}

func LinqMax[T constraints.Ordered](obj []T) T {
	if len(obj) == 0 {
		return Null[T]().ValueOrZero()
	}

	max := obj[0]

	for _, item := range obj {
		if item > max {
			max = item
		}
	}

	return max
}

func LinqMaxIterable[T constraints.Ordered, TV IIterable[T]](obj TV) T {
	max := Null[T]()
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if max.IsError() || item > max.Data {
			max = Some(item)
		}
	}

	return max.ValueOrZero()
}


func LinqMin[T constraints.Ordered](obj []T) T {
	if len(obj) == 0 {
		return Null[T]().ValueOrZero()
	}

	min := obj[0]

	for _, item := range obj {
		if item < min {
			min = item
		}
	}

	return min
}

func LinqMinIterable[T constraints.Ordered, TV IIterable[T]](obj TV) T {
	min := Null[T]()
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		if min.IsError() || item < min.Data {
			min = Some(item)
		}
	}

	return min.ValueOrZero()
}

func LinqReduce[T any, P any](obj []T, reducer func(T, P) P, acc P) P {
	for _, item := range obj {
		acc = reducer(item, acc)
	}

	return acc
}

func LinqReduceIterable[T any, TV IIterable[T], P any](obj TV, reducer func(T, P) P, acc P) P {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		acc = reducer(item, acc)
	}

	return acc
}

func LinqAggregate[T any, P any](obj []T, reducer func(T, P) P, acc P) P {
	return LinqReduce[T, P](obj, reducer, acc)
}

func LinqAggregateIterable[T any, TV IIterable[T], P any](obj TV, reducer func(T, P) P, acc P) P {
	return LinqReduceIterable[T, TV, P](obj, reducer, acc)
}

func LinqSum[T constraints.Ordered](obj []T) (sum T) {
	for _, item := range obj {
		sum += item
	}

	return sum
}

func LinqSumIterable[T constraints.Ordered, TV IIterable[T]](obj TV) (sum T) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		sum += item
	}

	return sum
}

func LinqSumBy[T any, R constraints.Ordered](obj []T, selector func(T) R) (sum R) {
	for _, item := range obj {
		sum += selector(item)
	}

	return sum
}

func LinqSumByIterable[T any, TV IIterable[T],R constraints.Ordered](obj TV, selector func(T) R) (sum R) {
	for {
		item, ok := obj.Next()
		if !ok {
			break
		}
		sum += selector(item)
	}

	return sum
}
