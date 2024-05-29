package defaultprovider

import (
	"github.com/patrykjadamczyk/go-utils/intl"
)

type DefaultTranslationProviderMessage struct {
	Msg       intl.IntlMessageStruct
	PluralMsg intl.IntlPluralMessageStruct
	GenderMsg intl.IntlGenderMessageStruct
}

type DefaultTranslationProvider struct {
	Messages map[intl.IntlLocale]map[intl.IntlTextDomain]map[intl.IntlMessageID]DefaultTranslationProviderMessage
	Locale   intl.IntlLocale
}

func (p *DefaultTranslationProvider) SetLocale(newLocale intl.IntlLocale) {
	p.Locale = newLocale
}

func (p *DefaultTranslationProvider) GetMessageUsingLocale(
	locale intl.IntlLocale,
	messageID intl.IntlMessageID,
) intl.IntlMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].Msg
}

func (p *DefaultTranslationProvider) GetMessageUsingLocaleAndDomain(
	locale intl.IntlLocale,
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlMessageStruct {
	return p.Messages[locale][domain][messageID].Msg
}

func (p *DefaultTranslationProvider) GetMessageUsingDomain(
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlMessageStruct {
	return p.Messages[p.Locale][domain][messageID].Msg
}

func (p *DefaultTranslationProvider) GetMessage(
	messageID intl.IntlMessageID,
) intl.IntlMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].Msg
}

func (p *DefaultTranslationProvider) GetPluralMessageUsingLocale(
	locale intl.IntlLocale,
	messageID intl.IntlMessageID,
) intl.IntlPluralMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].PluralMsg
}

func (p *DefaultTranslationProvider) GetPluralMessageUsingLocaleAndDomain(
	locale intl.IntlLocale,
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlPluralMessageStruct {
	return p.Messages[locale][domain][messageID].PluralMsg
}

func (p *DefaultTranslationProvider) GetPluralMessageUsingDomain(
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlPluralMessageStruct {
	return p.Messages[p.Locale][domain][messageID].PluralMsg
}

func (p *DefaultTranslationProvider) GetPluralMessage(
	messageID intl.IntlMessageID,
) intl.IntlPluralMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].PluralMsg
}

func (p *DefaultTranslationProvider) GetGenderMessageUsingLocale(
	locale intl.IntlLocale,
	messageID intl.IntlMessageID,
) intl.IntlGenderMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].GenderMsg
}

func (p *DefaultTranslationProvider) GetGenderMessageUsingLocaleAndDomain(
	locale intl.IntlLocale,
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlGenderMessageStruct {
	return p.Messages[locale][domain][messageID].GenderMsg
}

func (p *DefaultTranslationProvider) GetGenderMessageUsingDomain(
	domain intl.IntlTextDomain,
	messageID intl.IntlMessageID,
) intl.IntlGenderMessageStruct {
	return p.Messages[p.Locale][domain][messageID].GenderMsg
}

func (p *DefaultTranslationProvider) GetGenderMessage(
	messageID intl.IntlMessageID,
) intl.IntlGenderMessageStruct {
	return p.Messages[p.Locale][intl.INTL_DEFAULT_TEXT_DOMAIN][messageID].GenderMsg
}

func NewDefaultTranslationProvider() *DefaultTranslationProvider {
	return &DefaultTranslationProvider{
		Locale: intl.INTL_DEFAULT_LOCALE,
		Messages: make(
			map[intl.IntlLocale]map[intl.IntlTextDomain]map[intl.IntlMessageID]DefaultTranslationProviderMessage,
		),
	}
}

// func (p *DefaultTranslationProvider) Dev() {
// 	t := func(t intl.TranslationProvider) {
// 	}
// 	t(p)
// }
