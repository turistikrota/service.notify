package telegram

import (
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/telegram"
)

type repo struct {
	conf          config.Telegram
	factory       telegram.Factory
	notifyFactory notify.Factory
}

func New(factory telegram.Factory, notifyFactory notify.Factory, config config.Telegram) telegram.Repository {
	if notifyFactory.IsZero() {
		panic("notifyFactory is zero")
	}
	return &repo{
		conf:          config,
		factory:       factory,
		notifyFactory: notifyFactory,
	}
}
