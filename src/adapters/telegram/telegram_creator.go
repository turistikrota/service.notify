package telegram

import (
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/notify"
	"api.turistikrota.com/notify/src/domain/telegram"
)

type Creator interface {
	New(factory telegram.Factory, notifyFactory notify.Factory, config config.Telegram) telegram.Repository
}

type creator struct{}

func NewCreator() Creator {
	return &creator{}
}

func (c *creator) New(factory telegram.Factory, notifyFactory notify.Factory, config config.Telegram) telegram.Repository {
	return New(factory, notifyFactory, config)
}
