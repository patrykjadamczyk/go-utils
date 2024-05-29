package wrapper

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

func (w *Wrapper[T]) Len() int {
	return UniversalLen(w.Value)
}

// type IContains[T any] interface {
// 	Contains(T) bool
// }
//
// type IContainsAny interface {
// 	Contains(any) bool
// }
//
// type IContainsAnyFunc interface {
// 	ContainsAny(any) bool
// }
//
// type IInsert[T any] interface {
// 	Insert(T) T
// }
//
// type IInsertAny interface {
// 	Insert(any) any
// }
//
// type IInsertAnyFunc interface {
// 	InsertAny(any) any
// }
//
// type IInsertInplace[T any] interface {
// 	Insert(T)
// }
//
// type IInsertInplaceAny interface {
// 	Insert(any)
// }
//
// type IInsertInplaceAnyFunc interface {
// 	InsertAny(any)
// }
//
// type IAppend[T any] interface {
// 	Append(T) T
// }
//
// type IAppendAny interface {
// 	Append(any) any
// }
//
// type IAppendAnyFunc interface {
// 	AppendAny(any) any
// }
//
// type IAppendInplace[T any] interface {
// 	Append(T)
// }
//
// type IAppendInplaceAny interface {
// 	Append(any)
// }
//
// type IAppendInplaceAnyFunc interface {
// 	AppendAny(any)
// }
//
// type IPrepend[T any] interface {
// 	Prepend(T) T
// }
//
// type IPrependAny interface {
// 	Prepend(any) any
// }
//
// type IPrependAnyFunc interface {
// 	PrependAny(any) any
// }
//
// type IPrependInplace[T any] interface {
// 	Prepend(T)
// }
//
// type IPrependInplaceAny interface {
// 	Prepend(any)
// }
//
// type IPrependInplaceAnyFunc interface {
// 	PrependAny(any)
// }
