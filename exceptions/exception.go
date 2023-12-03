package exceptions

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/exceptions/categories"
)

type BasedException struct {
	WrappedError error
	SubCategoryExtendedException
}

func ConvertErrorToException(err error, subcategory ExceptionCategory) BasedException {
    res := BasedException{}
    res.Init(categories.GoBaseException, subcategory, err.Error())
    res.WrappedError = err
    return res
}
