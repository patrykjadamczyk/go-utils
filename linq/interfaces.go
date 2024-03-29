package linq

type IEnumerable[TV any] interface {
	*[]TV | *map[int]TV | *map[string]TV
}

type IEnumerableV[TK comparable, TV any] interface {
	*map[TK]TV
}

type IIterable[TV any] interface {
	Next() (TV, bool)
}

type IIterableV[TK comparable, TV any] interface {
	Next() (TK, TV, bool)
}

type IResetableIterable[TV any] interface {
	Reset()
	Next() (TV, bool)
}

type IResetableIterableV[TK comparable, TV any] interface {
	Reset()
	Next() (TK, TV, bool)
}
