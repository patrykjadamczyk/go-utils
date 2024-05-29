package wrapper

import (
	"reflect"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) ToArray() []T {
	val := reflect.ValueOf(w.Value)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return EnsureType[[]T](w.Value)
	}
	return []T{w.Value}
}

func (w *Wrapper[T]) ToArrayAny() (res []any) {
	val := reflect.ValueOf(w.Value)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		res = make([]any, 0)
		for _, item := range EnsureType[[]T](w.Value) {
            res = append(res, item)
        }
		return
	}
	return []any{w.Value}
}
