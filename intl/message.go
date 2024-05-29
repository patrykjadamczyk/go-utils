package intl

import (
	"fmt"

	"github.com/patrykjadamczyk/go-utils/intl/gender"
)

// Message Structure
type IntlMessageStruct struct {
	Message IntlMessage
}

// Plural Message Structure
type IntlPluralMessageStruct struct {
	PluralMessage func(any) IntlMessage
	Message       IntlMessage
}

// Message with gender Structure
type IntlGenderMessageStruct struct {
	He      IntlMessage
	She     IntlMessage
	They    IntlMessage
	Neutral IntlMessage
}

func (im IntlMessageStruct) Get() IntlMessage {
	return im.Message
}

func (im IntlMessageStruct) Format(args ...any) IntlMessage {
	return IntlMessage(fmt.Sprintf(string(im.Get()), args...))
}

func (imp IntlPluralMessageStruct) Get() IntlMessage {
	return imp.Message
}

func (imp IntlPluralMessageStruct) GetN(n any) IntlMessage {
	return imp.PluralMessage(n)
}

func (imp IntlPluralMessageStruct) Format(args ...any) IntlMessage {
	return IntlMessage(fmt.Sprintf(string(imp.Get()), args...))
}

func (imp IntlPluralMessageStruct) FormatN(n any, args ...any) IntlMessage {
	return IntlMessage(fmt.Sprintf(string(imp.GetN(n)), args...))
}

func (imp IntlGenderMessageStruct) Get() IntlMessage {
	return imp.Neutral
}

func (igms IntlGenderMessageStruct) GetG(g gender.IntlGender) IntlMessage {
	switch g {
	case gender.He:
		return igms.He
	case gender.She:
		return igms.She
	case gender.They:
		return igms.They
	case gender.Neutral:
		return igms.Neutral
	default:
		return igms.Neutral
	}
}

func (imp IntlGenderMessageStruct) Format(args ...any) IntlMessage {
	return IntlMessage(fmt.Sprintf(string(imp.Get()), args...))
}

func (imp IntlGenderMessageStruct) FormatG(g gender.IntlGender, args ...any) IntlMessage {
	return IntlMessage(fmt.Sprintf(string(imp.GetG(g)), args...))
}

type IntlMessageInterface interface {
	Get() IntlMessage

	Format() IntlMessage
}

type IntlPluralMessageInterface interface {
	Get() IntlMessage
	GetN(n any) IntlMessage

	Format() IntlMessage
	FormatN(any) IntlMessage
}

type IntlGenderMessageInterface interface {
	Get() IntlMessage
	GetG(g gender.IntlGender) IntlMessage

	Format() IntlMessage
	FormatN(g gender.IntlGender) IntlMessage
}
