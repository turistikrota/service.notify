package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/push"
)

type SendPushCommand struct {
	Notification *push.Notification
	Token        string
}

type SendPushResult struct{}

type SendPushHandler decorator.CommandHandler[SendPushCommand, *SendPushResult]

type sendPushHandler struct {
	pushRepo      push.Repository
	notifyRepo    notify.Repository
	pushFactory   push.Factory
	notifyFactory notify.Factory
}

type SendPushHandlerConfig struct {
	PushRepo      push.Repository
	NotifyRepo    notify.Repository
	PushFactory   push.Factory
	NotifyFactory notify.Factory
	CqrsBase      decorator.Base
}

func NewSendPushHandler(config SendPushHandlerConfig) SendPushHandler {
	return decorator.ApplyCommandDecorators[SendPushCommand, *SendPushResult](
		sendPushHandler{
			pushRepo:      config.PushRepo,
			notifyRepo:    config.NotifyRepo,
			pushFactory:   config.PushFactory,
			notifyFactory: config.NotifyFactory,
		},
		config.CqrsBase,
	)
}

func (h sendPushHandler) Handle(ctx context.Context, command SendPushCommand) (*SendPushResult, *i18np.Error) {
	m, err := h.pushFactory.NewNotifyPush(command.Token, command.Notification)
	if err != nil {
		return nil, err
	}
	err = h.pushFactory.Validate(m.Notify)
	if err != nil {
		return nil, err
	}
	err = h.pushRepo.Send(ctx, m, m.Token)
	if err != nil {
		return nil, err
	}
	err = h.notifyRepo.Log(ctx, m.Notify, m.Data)
	if err != nil {
		return nil, err
	}
	return &SendPushResult{}, nil
}
