package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.notify/src/domain/mail"
	"github.com/turistikrota/service.notify/src/domain/notify"
)

type SendMailCommand struct {
	Recipient string
	Data      *mail.Data
}

type SendMailResult struct{}

type SendMailHandler decorator.CommandHandler[SendMailCommand, *SendMailResult]

type sendMailHandler struct {
	mailRepo      mail.Repository
	notifyRepo    notify.Repository
	mailFactory   mail.Factory
	notifyFactory notify.Factory
}

type SendMailHandlerConfig struct {
	MailRepo      mail.Repository
	NotifyRepo    notify.Repository
	MailFactory   mail.Factory
	NotifyFactory notify.Factory
	CqrsBase      decorator.Base
}

func NewSendMailHandler(config SendMailHandlerConfig) SendMailHandler {
	return decorator.ApplyCommandDecorators[SendMailCommand, *SendMailResult](
		sendMailHandler{
			mailRepo:      config.MailRepo,
			notifyRepo:    config.NotifyRepo,
			mailFactory:   config.MailFactory,
			notifyFactory: config.NotifyFactory,
		},
		config.CqrsBase,
	)
}

func (h sendMailHandler) Handle(ctx context.Context, command SendMailCommand) (*SendMailResult, *i18np.Error) {
	m, err := h.mailFactory.NewNotifyMail(command.Recipient, command.Data)
	if err != nil {
		return nil, err
	}
	err = h.mailFactory.Validate(m.Notify)
	if err != nil {
		return nil, err
	}
	err = h.mailRepo.Send(ctx, mail.SendConfig{
		Recipient: command.Recipient,
		Data:      command.Data,
	})
	if err != nil {
		return nil, err
	}
	err = h.notifyRepo.Log(ctx, m.Notify, m.Data)
	if err != nil {
		return nil, err
	}
	return &SendMailResult{}, nil
}
