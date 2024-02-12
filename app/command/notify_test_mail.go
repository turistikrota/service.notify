package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/mail"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifyTestMailCmd struct {
	Email  string `json:"email" validate:"required,email"`
	Locale string `json:"-"`
}

type NotifyTestMailRes struct{}

type NotifyTestMailHandler cqrs.HandlerFunc[NotifyTestMailCmd, *NotifyTestMailRes]

func NewNotifyTestMailHandler(factory notify.Factory, i18n *i18np.I18n, srv mail.Service) NotifyTestMailHandler {
	return func(ctx context.Context, cmd NotifyTestMailCmd) (*NotifyTestMailRes, *i18np.Error) {
		err := srv.SendText(mail.SendConfig{
			To:      cmd.Email,
			Subject: i18n.Translate(factory.Messages.TestMailSubject, cmd.Locale),
			Message: i18n.Translate(factory.Messages.TestMailContent, cmd.Locale),
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &NotifyTestMailRes{}, nil
	}
}
