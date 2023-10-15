package notify

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	Failed(operation string) *i18np.Error
	TypeNotFound(t string) *i18np.Error
	ValidationFailed(operation string) *i18np.Error
	NotFound() *i18np.Error
}

type notifyErrors struct{}

func newNotifyErrors() Errors {
	return &notifyErrors{}
}

func (e *notifyErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *notifyErrors) TypeNotFound(t string) *i18np.Error {
	return i18np.NewError(I18nMessages.TypeNotFound, i18np.P{"Type": t})
}

func (e *notifyErrors) ValidationFailed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": operation})
}

func (e *notifyErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}
