package linq

func LinqWhere[T any](obj []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(obj))

	for _, item := range obj {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func LinqWhereIterable[T any, TV IIterable[T]](obj TV, predicate func(T) bool) []T {
    result := make([]T, 0)

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

func LinqSelect[IT any, OT any](obj []IT, predicate func(IT) OT) []OT {
    result := make([]OT, 0, len(obj))

    for _, item := range obj {
        result = append(result, predicate(item))
    }

    return result
}

func LinqSelectIterable[IT any, TV IIterable[IT], OT any](obj TV, predicate func(IT) OT) []OT {
    result := make([]OT, 0)

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

func LinqTakeIterable[T any, TV IIterable[T]](obj TV, count int) []T {
    result := make([]T, 0)

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

func LinqSkipIterable[T any, TV IIterable[T]](obj TV, count int) []T {
    result := make([]T, 0)
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
