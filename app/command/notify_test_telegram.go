package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/telegram"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifyTestTelegramCmd struct {
	ChatID string `json:"chat_id" validate:"required"`
	Locale string `json:"-"`
}

type NotifyTestTelegramRes struct{}

type NotifyTestTelegramHandler cqrs.HandlerFunc[NotifyTestTelegramCmd, *NotifyTestTelegramRes]

func NewNotifyTestTelegramHandler(factory notify.Factory, i18n *i18np.I18n, srv telegram.Service) NotifyTestTelegramHandler {
	return func(ctx context.Context, cmd NotifyTestTelegramCmd) (*NotifyTestTelegramRes, *i18np.Error) {
		err := srv.Send(ctx, telegram.SendConfig{
			ChatID: cmd.ChatID,
			Text:   i18n.Translate(factory.Messages.TestTelegramContent, cmd.Locale),
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &NotifyTestTelegramRes{}, nil
	}
}
