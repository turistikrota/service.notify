package sms

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	Failed(operation string) *i18np.Error
	ValidationFailed(operation string) *i18np.Error
}

type smsErrors struct{}

func newSmsErrors() Errors {
	return &smsErrors{}
}

func (e *smsErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *smsErrors) ValidationFailed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": operation})
}
