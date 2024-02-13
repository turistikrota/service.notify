package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/mail"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifySendEmailCmd struct {
	ActorName      string `json:"actorName"`
	Template       string `json:"template"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	TemplateParams any    `json:"templateParams"`
	Translate      bool   `json:"translate"`
	Locale         string `json:"locale"`
}

type NotifySendEmailRes struct{}

type NotifySendEmailHandler cqrs.HandlerFunc[NotifySendEmailCmd, *NotifySendEmailRes]

func NewNotifySendEmailHandler(factory notify.Factory, actorConfigRepo actor_config.Repository, i18n *i18np.I18n, srv mail.Service) NotifySendEmailHandler {

	sender := func(cmd NotifySendEmailCmd, email string) error {
		subject := cmd.Subject
		content := cmd.Content
		if cmd.Translate {
			subject = i18n.Translate(subject, cmd.Locale)
			content = i18n.Translate(content, cmd.Locale)
		}
		if cmd.Template != "" {
			return srv.SendWithTemplate(mail.SendWithTemplateConfig{
				SendConfig: mail.SendConfig{
					To:      email,
					Subject: subject,
					Message: content,
				},
				Template: cmd.Template,
				Data:     cmd.TemplateParams,
			})
		}
		return srv.SendText(mail.SendConfig{
			To:      email,
			Subject: subject,
			Message: content,
		})
	}

	return func(ctx context.Context, cmd NotifySendEmailCmd) (*NotifySendEmailRes, *i18np.Error) {
		config, err := actorConfigRepo.GetByActorName(ctx, cmd.ActorName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		if len(config.Mail) == 0 {
			return nil, factory.Errors.NotMailConfigured()
		}
		for _, mailConfig := range config.Mail {
			err := sender(cmd, mailConfig.Email)
			if err != nil {
				return nil, factory.Errors.Failed(err.Error())
			}
		}
		return &NotifySendEmailRes{}, nil
	}
}
