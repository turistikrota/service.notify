package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.notify/adapters/mail"
	"github.com/turistikrota/service.notify/adapters/sms"
	"github.com/turistikrota/service.notify/adapters/telegram"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/domains/notify"
)

type NotifySendToAllChannelsCmd struct {
	ActorName    string `json:"actorName"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Locale       string `json:"locale"`
	Template     string `json:"template,omitempty"`
	TemplateData any    `json:"templateData,omitempty"`
	Translate    bool   `json:"translate"`
}

type NotifySendToAllChannelsRes struct{}

type NotifySendToAllChannelsHandler cqrs.HandlerFunc[NotifySendToAllChannelsCmd, *NotifySendToAllChannelsRes]

func NewNotifySendToAllChannelsHandler(factory notify.Factory, actorConfigRepo actor_config.Repository, i18n *i18np.I18n, smsSrv sms.Service, mailSrv mail.Service, telegramSrv telegram.Service) NotifySendToAllChannelsHandler {
	return func(ctx context.Context, cmd NotifySendToAllChannelsCmd) (*NotifySendToAllChannelsRes, *i18np.Error) {
		config, err := actorConfigRepo.GetByActorName(ctx, cmd.ActorName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		subject := cmd.Subject
		content := cmd.Content
		if cmd.Translate {
			subject = i18n.Translate(subject, cmd.Locale)
			content = i18n.Translate(content, cmd.Locale)
		}
		go func() {
			for _, cnf := range config.SMS {
				err := smsSrv.Send(ctx, sms.SendConfig{
					Phone: cnf.CountryCode + cnf.Phone,
					Text:  content,
					Lang:  cmd.Locale,
				})
				if err != nil {
					return
				}
			}
		}()
		go func() {
			for _, cnf := range config.Mail {
				var err error
				if cmd.Template != "" {
					err = mailSrv.SendWithTemplate(mail.SendWithTemplateConfig{
						SendConfig: mail.SendConfig{
							To:      cnf.Email,
							Subject: subject,
							Message: content,
						},
						Template: cmd.Template,
						Data:     cmd.TemplateData,
					})
				} else {
					err = mailSrv.SendText(mail.SendConfig{
						To:      cnf.Email,
						Subject: subject,
						Message: content,
					})
				}
				if err != nil {
					return
				}
			}
		}()
		go func() {
			for _, cnf := range config.Telegram {
				err := telegramSrv.Send(ctx, telegram.SendConfig{
					ChatID: cnf.ChatID,
					Text:   content,
				})
				if err != nil {
					return
				}
			}
		}()
		return &NotifySendToAllChannelsRes{}, nil
	}
}
