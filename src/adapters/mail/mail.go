package mail

import (
	"github.com/turistikrota/service.notify/src/adapters/mail/google"
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/mail"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type Mail interface {
	NewGoogle(factory mail.Factory, notifyFactory notify.Factory, config config.MailGoogle) mail.Repository
}

type mailClient struct{}

func New() Mail {
	return &mailClient{}
}

func (m *mailClient) NewGoogle(factory mail.Factory, notifyFactory notify.Factory, config config.MailGoogle) mail.Repository {
	return google.New(factory, notifyFactory, config)
}
