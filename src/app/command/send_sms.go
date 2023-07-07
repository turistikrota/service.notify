package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/sms"
)

type SendSmsCommand struct {
	Recipient string
	Data      *sms.Data
}

type SendSmsResult struct{}

type SendSmsHandler decorator.CommandHandler[SendSmsCommand, *SendSmsResult]

type sendSmsHandler struct {
	smsRepo       sms.Repository
	notifyRepo    notify.Repository
	smsFactory    sms.Factory
	notifyFactory notify.Factory
}

type SendSmsHandlerConfig struct {
	SmsRepo       sms.Repository
	NotifyRepo    notify.Repository
	SmsFactory    sms.Factory
	NotifyFactory notify.Factory
	CqrsBase      decorator.Base
}

func NewSendSmsHandler(config SendSmsHandlerConfig) SendSmsHandler {
	return decorator.ApplyCommandDecorators[SendSmsCommand, *SendSmsResult](
		sendSmsHandler{
			smsRepo:       config.SmsRepo,
			notifyRepo:    config.NotifyRepo,
			smsFactory:    config.SmsFactory,
			notifyFactory: config.NotifyFactory,
		},
		config.CqrsBase,
	)
}

func (h sendSmsHandler) Handle(ctx context.Context, command SendSmsCommand) (*SendSmsResult, *i18np.Error) {
	m, err := h.smsFactory.NewNotifySms(command.Recipient, command.Data)
	if err != nil {
		return nil, err
	}
	err = h.smsFactory.Validate(m.Notify)
	if err != nil {
		return nil, err
	}
	err = h.smsRepo.Send(ctx, sms.SendConfig{
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
	return &SendSmsResult{}, nil
}
