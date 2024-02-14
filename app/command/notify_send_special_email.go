package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/mail"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifySendSpecialEmailCmd struct {
	Email          string `json:"email"`
	Template       string `json:"template"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	TemplateParams any    `json:"templateParams"`
	Translate      bool   `json:"translate"`
	Locale         string `json:"locale"`
}

type NotifySendSpecialEmailRes struct{}

type NotifySendSpecialEmailHandler cqrs.HandlerFunc[NotifySendSpecialEmailCmd, *NotifySendSpecialEmailRes]

func NewNotifySendSpecialEmailHandler(factory notify.Factory, i18n *i18np.I18n, srv mail.Service) NotifySendSpecialEmailHandler {
	return func(ctx context.Context, cmd NotifySendSpecialEmailCmd) (*NotifySendSpecialEmailRes, *i18np.Error) {
		subject := cmd.Subject
		content := cmd.Content
		if cmd.Translate {
			subject = i18n.Translate(subject, cmd.Locale)
			content = i18n.Translate(content, cmd.Locale)
		}
		if cmd.Template != "" {
			go srv.SendWithTemplate(mail.SendWithTemplateConfig{
				SendConfig: mail.SendConfig{
					To:      cmd.Email,
					Subject: subject,
					Message: content,
				},
				Template: cmd.Template,
				Data:     cmd.TemplateParams,
			})
		} else {
			go srv.SendText(mail.SendConfig{
				To:      cmd.Email,
				Subject: subject,
				Message: content,
			})
		}
		return &NotifySendSpecialEmailRes{}, nil
	}
}
