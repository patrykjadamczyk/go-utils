package interfaces

type IToTuple2[T1, T2 any] interface {
    ToTuple() (T1, T2)
}

type IToTuple3[T1, T2, T3 any] interface {
    ToTuple() (T1, T2, T3)
}

type IToTuple4[T1, T2, T3, T4 any] interface {
    ToTuple() (T1, T2, T3, T4)
}
