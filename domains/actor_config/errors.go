package actor_config

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	NotFound() *i18np.Error
	InvalidUUID() *i18np.Error
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

func (e *errors) NotFound() *i18np.Error {
	return i18np.NewError(i18nMessages.NotFound)
}

func (e *errors) InvalidUUID() *i18np.Error {
	return i18np.NewError(i18nMessages.InvalidUUID)
}
