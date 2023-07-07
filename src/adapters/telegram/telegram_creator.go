package telegram

import (
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/telegram"
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
