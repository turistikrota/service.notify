package google

import (
	"net/smtp"

	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/mail"
	"api.turistikrota.com/notify/src/domain/notify"
	"github.com/sirupsen/logrus"
)

type repo struct {
	conf          config.MailGoogle
	auth          smtp.Auth
	factory       mail.Factory
	notifyFactory notify.Factory
}

func New(factory mail.Factory, notifyFactory notify.Factory, config config.MailGoogle) mail.Repository {
	if notifyFactory.IsZero() {
		panic("notifyFactory is zero")
	}
	r := &repo{
		conf:          config,
		factory:       factory,
		notifyFactory: notifyFactory,
	}
	return r.authorize()
}

func (r *repo) authorize() *repo {
	logrus.Info("With username and pw", r.conf.Username, " : ", r.conf.Password)
	r.auth = smtp.PlainAuth(r.conf.Identity, r.conf.Username, r.conf.Password, r.conf.SmtpHost)
	return r
}
