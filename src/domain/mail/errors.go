package mail

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	Failed(operation string) *i18np.Error
	ValidationFailed(operation string) *i18np.Error
}

type mailErrors struct{}

func newSmsErrors() Errors {
	return &mailErrors{}
}

func (e *mailErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *mailErrors) ValidationFailed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": operation})
}
