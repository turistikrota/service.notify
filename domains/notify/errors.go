package notify

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	NotMailConfigured() *i18np.Error
	NotSmsConfigured() *i18np.Error
}

type errors struct{}

func newErrors() Errors {
	return &errors{}
}

func (e *errors) Failed(operation string) *i18np.Error {
	return i18np.NewError(i18nMessages.Failed, i18np.P{
		"Operation": operation,
	})
}

func (e *errors) NotMailConfigured() *i18np.Error {
	return i18np.NewError(i18nMessages.NotMailConfigured)
}

func (e *errors) NotSmsConfigured() *i18np.Error {
	return i18np.NewError(i18nMessages.NotSmsConfigured)
}
