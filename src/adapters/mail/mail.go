package mail

import (
	"api.turistikrota.com/notify/src/adapters/mail/google"
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/mail"
	"api.turistikrota.com/notify/src/domain/notify"
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
