package telegram

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	Failed(operation string) *i18np.Error
	ValidationFailed(operation string) *i18np.Error
}

type telegramErrors struct{}

func newSmsErrors() Errors {
	return &telegramErrors{}
}

func (e *telegramErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *telegramErrors) ValidationFailed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": operation})
}
