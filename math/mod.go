package math

func Mod[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~int | ~uint](
	x T,
	y T,
) T {
	var v T
	value := any(v)
	switch value.(type) {
	case float64:
		return ensureType[T](Fmod(ensureType[float64](x), ensureType[float64](y)))
	case float32:
		return ensureType[T](Fmod(ensureType[float32](x), ensureType[float32](y)))
	case int:
		return ensureType[T](ensureType[int](x) % ensureType[int](y))
	case int8:
		return ensureType[T](ensureType[int8](x) % ensureType[int8](y))
	case int16:
		return ensureType[T](ensureType[int16](x) % ensureType[int16](y))
	case int32:
		return ensureType[T](ensureType[int32](x) % ensureType[int32](y))
	case int64:
		return ensureType[T](ensureType[int64](x) % ensureType[int64](y))
	case uint:
		return ensureType[T](ensureType[uint](x) % ensureType[uint](y))
	case uint8:
		return ensureType[T](ensureType[uint8](x) % ensureType[uint8](y))
	case uint16:
		return ensureType[T](ensureType[uint16](x) % ensureType[uint16](y))
	case uint32:
		return ensureType[T](ensureType[uint32](x) % ensureType[uint32](y))
	case uint64:
		return ensureType[T](ensureType[uint64](x) % ensureType[uint64](y))
	}
	return T(0)
}
