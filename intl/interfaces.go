package intl

type TranslationProvider interface {
	GetMessageUsingLocale(locale IntlLocale, messageID IntlMessageID) IntlMessageStruct
	GetMessageUsingLocaleAndDomain(locale IntlLocale, domain IntlTextDomain, messageID IntlMessageID) IntlMessageStruct
	GetMessageUsingDomain(domain IntlTextDomain, messageID IntlMessageID) IntlMessageStruct
	GetMessage(messageID IntlMessageID) IntlMessageStruct

	GetPluralMessageUsingLocale(locale IntlLocale, messageID IntlMessageID) IntlPluralMessageStruct
	GetPluralMessageUsingLocaleAndDomain(locale IntlLocale, domain IntlTextDomain, messageID IntlMessageID) IntlPluralMessageStruct
	GetPluralMessageUsingDomain(domain IntlTextDomain, messageID IntlMessageID) IntlPluralMessageStruct
	GetPluralMessage(messageID IntlMessageID) IntlPluralMessageStruct

	GetGenderMessageUsingLocale(locale IntlLocale, messageID IntlMessageID) IntlGenderMessageStruct
	GetGenderMessageUsingLocaleAndDomain(locale IntlLocale, domain IntlTextDomain, messageID IntlMessageID) IntlGenderMessageStruct
	GetGenderMessageUsingDomain(domain IntlTextDomain, messageID IntlMessageID) IntlGenderMessageStruct
	GetGenderMessage(messageID IntlMessageID) IntlGenderMessageStruct

	SetLocale(locale IntlLocale)
}
