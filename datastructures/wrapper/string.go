package wrapper

import (
	"fmt"

	"github.com/patrykjadamczyk/go-utils/datastructures/interfaces"
)

func (w *Wrapper[T]) String() string {
	if v, ok := any(w.Value).(fmt.Stringer); ok {
        return v.String()
    }
	if v, ok := any(w.Value).(fmt.GoStringer); ok {
        return v.GoString()
    }
	if v, ok := any(w.Value).(interfaces.IString); ok {
		return v.String()
	}
	if v, ok := any(w.Value).(interfaces.IToString); ok {
		return v.ToString()
	}
	wv := w.Get()
    return fmt.Sprintf("%T(%v)", wv, wv)
}

func (w *Wrapper[T]) ToString() string {
    return w.String()
}

func (w *Wrapper[T]) GoString() string {
	return w.String()
}

func (w *Wrapper[T]) FromString(s string) {
	if v, ok := any(w.Value).(interfaces.IFromString); ok {
        v.FromString(s)
    } else {
		panic("Value is not IFromString")
	}
}
