package wrapper

import (
	"reflect"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) ToArray() []T {
	wv := w.Get()
	val := reflect.ValueOf(wv)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return EnsureType[[]T](wv)
	}
	return []T{wv}
}

func (w *Wrapper[T]) ToArrayAny() (res []any) {
	wv := w.Get()
	val := reflect.ValueOf(wv)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		res = make([]any, 0)
		for _, item := range EnsureType[[]T](wv) {
			res = append(res, item)
		}
		return
	}
	return []any{wv}
}
