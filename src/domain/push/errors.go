package push

import "github.com/mixarchitecture/i18np"

type Errors interface {
	Failed(operation string) *i18np.Error
	ValidationFailed(operation string) *i18np.Error
}

type pushErrors struct{}

func newPushErrors() Errors {
	return &pushErrors{}
}

func (e *pushErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *pushErrors) ValidationFailed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.ValidationFailed, i18np.P{"Operation": operation})
}
