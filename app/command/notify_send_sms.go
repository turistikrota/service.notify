package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/sms"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifySendSmsCmd struct {
	ActorName string `json:"actorName"`
	Text      string `json:"text"`
	Locale    string `json:"locale"`
	Translate bool   `json:"translate"`
}

type NotifySendSmsRes struct{}

type NotifySendSmsHandler cqrs.HandlerFunc[NotifySendSmsCmd, *NotifySendSmsRes]

func NewNotifySendSmsHandler(factory notify.Factory, actorConfigRepo actor_config.Repository, i18n *i18np.I18n, srv sms.Service) NotifySendSmsHandler {
	return func(ctx context.Context, cmd NotifySendSmsCmd) (*NotifySendSmsRes, *i18np.Error) {
		config, err := actorConfigRepo.GetByActorName(ctx, cmd.ActorName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		if len(config.SMS) == 0 {
			return nil, factory.Errors.NotSmsConfigured()
		}
		txt := cmd.Text
		if cmd.Translate {
			txt = i18n.Translate(txt, cmd.Locale)
		}
		for _, phone := range config.SMS {
			err := srv.Send(ctx, sms.SendConfig{
				Phone: phone.CountryCode + phone.Phone,
				Text:  txt,
				Lang:  cmd.Locale,
			})
			if err != nil {
				return nil, factory.Errors.Failed(err.Error())
			}
		}
		return &NotifySendSmsRes{}, nil
	}
}
