package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/telegram"
	"github.com/turistikrota/service.shared/chain"
	"github.com/turistikrota/service.shared/decorator"
)

type SendTelegramCommand struct {
	Recipient string
	Data      *telegram.Data
}

type SendTelegramResult struct{}

type SendTelegramHandler decorator.CommandHandler[SendTelegramCommand, *SendTelegramResult]

type sendTelegramHandler struct {
	telegramRepo    telegram.Repository
	notifyRepo      notify.Repository
	telegramFactory telegram.Factory
	notifyFactory   notify.Factory
}

type SendTelegramHandlerConfig struct {
	TelegramRepo    telegram.Repository
	NotifyRepo      notify.Repository
	TelegramFactory telegram.Factory
	NotifyFactory   notify.Factory
	CqrsBase        decorator.Base
}

type chainConfig struct {
	Command  SendTelegramCommand
	Telegram *telegram.Telegram
}

func NewSendTelegramHandler(config SendTelegramHandlerConfig) SendTelegramHandler {
	return decorator.ApplyCommandDecorators[SendTelegramCommand, *SendTelegramResult](
		sendTelegramHandler{
			telegramRepo:    config.TelegramRepo,
			notifyRepo:      config.NotifyRepo,
			telegramFactory: config.TelegramFactory,
			notifyFactory:   config.NotifyFactory,
		},
		config.CqrsBase,
	)
}

func (h sendTelegramHandler) Handle(ctx context.Context, command SendTelegramCommand) (*SendTelegramResult, *i18np.Error) {
	m, err := h.telegramFactory.NewNotifyTelegram(command.Recipient, command.Data)
	ch := chain.Make[chainConfig, SendTelegramResult]()
	ch.Use(h.validate, h.send, h.log, h.end)
	return ch.StartWithErr(ctx, chainConfig{
		Command:  command,
		Telegram: m,
	}, err)
}

func (h sendTelegramHandler) validate(ctx context.Context, config chainConfig) (*SendTelegramResult, *i18np.Error) {
	return nil, h.telegramFactory.Validate(config.Telegram.Notify)
}

func (h sendTelegramHandler) send(ctx context.Context, config chainConfig) (*SendTelegramResult, *i18np.Error) {
	return nil, h.telegramRepo.Send(ctx, telegram.SendConfig{
		Recipient: config.Command.Recipient,
		Data:      config.Command.Data,
	})
}

func (h sendTelegramHandler) log(ctx context.Context, config chainConfig) (*SendTelegramResult, *i18np.Error) {
	return nil, h.notifyRepo.Log(ctx, config.Telegram.Notify, config.Telegram.Data)
}

func (h sendTelegramHandler) end(ctx context.Context, config chainConfig) (*SendTelegramResult, *i18np.Error) {
	return &SendTelegramResult{}, nil
}
