package exceptions

import (
	"strings"

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

type switchOnException struct {
	SwitchValue Exception
}

func (e switchOnException) Case(category ExceptionCategory, f func()) switchOnException {
	if e.SwitchValue.Category == category {
		f()
	}
	scat := strings.Split(string(e.SwitchValue.Category), ":")
	if len(scat) >= 1 && scat[0] == string(category) {
		f()
	}

	return e
}

func (e switchOnException) CaseSubcategory(subcategory ExceptionCategory, f func()) switchOnException {
	scat := strings.Split(string(e.SwitchValue.Category), ":")
	if len(scat) >= 1 && scat[1] == string(subcategory) {
		f()
	}

	return e
}

func SwitchOnException(err Exception) switchOnException {
	return switchOnException{SwitchValue: err}
}
