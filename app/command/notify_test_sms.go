package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/sms"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifyTestSmsCmd struct {
	Phone  string `json:"phone" validate:"required,e164"`
	Locale string `json:"-"`
}

type NotifyTestSmsRes struct{}

type NotifyTestSmsHandler cqrs.HandlerFunc[NotifyTestSmsCmd, *NotifyTestSmsRes]

func NewNotifyTestSmsHandler(factory notify.Factory, i18n *i18np.I18n, srv sms.Service) NotifyTestSmsHandler {
	return func(ctx context.Context, cmd NotifyTestSmsCmd) (*NotifyTestSmsRes, *i18np.Error) {
		err := srv.Send(ctx, sms.SendConfig{
			Phone: cmd.Phone,
			Text:  i18n.Translate(factory.Messages.TestSmsContent, cmd.Locale),
			Lang:  cmd.Locale,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &NotifyTestSmsRes{}, nil
	}
}
