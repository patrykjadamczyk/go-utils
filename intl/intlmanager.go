package intl

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"github.com/patrykjadamczyk/go-utils/intl/gender"
)

type intlManager struct {
	Provider      TranslationProvider
	lastDomain    IntlTextDomain
	CurrentLocale Observable[IntlLocale]
	CurrentDomain Observable[IntlTextDomain]
}

// Set current domain of messages
func (im *intlManager) SetDomain(newDomain IntlTextDomain) {
	im.lastDomain = im.CurrentDomain.GetValue()
	im.CurrentDomain.SetValue(newDomain)
}

// Get current domain of messages
func (im *intlManager) GetDomain() IntlTextDomain {
	return im.CurrentDomain.GetValue()
}

// Restore domain of messages
// This should be used like that:
// IntlManager.SetDomain("Test")
// defer IntlManager.RestoreDomain()
func (im *intlManager) RestoreDomain() {
	im.CurrentDomain.SetValue(im.lastDomain)
}

// Restore domain of messages to default domain
// This should be used in situations where in one main scope you need additional change of domain
func (im *intlManager) RestoreDefaultDomain() {
	im.CurrentDomain.SetValue(INTL_DEFAULT_TEXT_DOMAIN)
}

// Set current locale
func (im *intlManager) SetLocale(newLocale IntlLocale) {
	im.CurrentLocale.SetValue(newLocale)
}

// Set default locale this function is pretty much the same as set current locale but avoids calling any observable handlers
func (im *intlManager) SetDefaultLocale(newLocale IntlLocale) {
	im.CurrentLocale.Value = newLocale
}

// Get current locale
func (im *intlManager) GetLocale() IntlLocale {
	return im.CurrentLocale.GetValue()
}

// Set translation provider
func (im *intlManager) SetTranslationProvider(provider TranslationProvider) {
	im.Provider = provider
}

// Get translation provider
func (im *intlManager) getTranslationProvider() TranslationProvider {
	AssertCustomError(im.Provider != nil, NewError("Translation provider not set"))
	return im.Provider
}

// Translate specified message id
func (im *intlManager) Translate(key IntlMessageID) IntlMessage {
	return im.
		getTranslationProvider().
		GetMessageUsingLocaleAndDomain(im.GetLocale(), im.GetDomain(), key).
		Get()
}

// Translate specified message id into plural message
func (im *intlManager) TranslateN(key IntlMessageID, n any) IntlMessage {
	return im.
		getTranslationProvider().
		GetPluralMessageUsingLocaleAndDomain(im.GetLocale(), im.GetDomain(), key).
		GetN(n)
}

// Translate specified message id into genderized message
func (im *intlManager) TranslateG(key IntlMessageID, g gender.IntlGender) IntlMessage {
	return im.
		getTranslationProvider().
		GetGenderMessageUsingLocaleAndDomain(im.GetLocale(), im.GetDomain(), key).
		GetG(g)
}

// Translate specified message id
func (im *intlManager) TranslateD(domain IntlTextDomain, key IntlMessageID) IntlMessage {
	return im.
		getTranslationProvider().
		GetMessageUsingLocaleAndDomain(im.GetLocale(), domain, key).
		Get()
}

// Translate specified message id into plural message
func (im *intlManager) TranslateDN(domain IntlTextDomain, key IntlMessageID, n any) IntlMessage {
	return im.
		getTranslationProvider().
		GetPluralMessageUsingLocaleAndDomain(im.GetLocale(), domain, key).
		GetN(n)
}

// Translate specified message id into genderized message
func (im *intlManager) TranslateDG(domain IntlTextDomain, key IntlMessageID, g gender.IntlGender) IntlMessage {
	return im.
		getTranslationProvider().
		GetGenderMessageUsingLocaleAndDomain(im.GetLocale(), domain, key).
		GetG(g)
}

var IntlManager = intlManager{
	CurrentLocale: MakeObservable(INTL_DEFAULT_LOCALE),
	CurrentDomain: MakeObservable(INTL_DEFAULT_TEXT_DOMAIN),
	lastDomain:    INTL_DEFAULT_TEXT_DOMAIN,
}
