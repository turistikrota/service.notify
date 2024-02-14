package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/sms"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifySendSpecialSmsCmd struct {
	Phone     string `json:"phone"`
	Text      string `json:"text"`
	Locale    string `json:"locale"`
	Translate bool   `json:"translate"`
}

type NotifySendSpecialSmsRes struct{}

type NotifySendSpecialSmsHandler cqrs.HandlerFunc[NotifySendSpecialSmsCmd, *NotifySendSpecialSmsRes]

func NewNotifySendSpecialSmsHandler(factory notify.Factory, i18n *i18np.I18n, srv sms.Service) NotifySendSpecialSmsHandler {
	return func(ctx context.Context, cmd NotifySendSpecialSmsCmd) (*NotifySendSpecialSmsRes, *i18np.Error) {
		txt := cmd.Text
		if cmd.Translate {
			txt = i18n.Translate(txt, cmd.Locale)
		}
		go srv.Send(ctx, sms.SendConfig{
			Phone: cmd.Phone,
			Text:  txt,
			Lang:  cmd.Locale,
		})
		return &NotifySendSpecialSmsRes{}, nil
	}
}
