package telegram

import (
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/notify"
	"api.turistikrota.com/notify/src/domain/telegram"
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
