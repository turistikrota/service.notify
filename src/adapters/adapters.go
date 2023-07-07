package adapters

import (
	"github.com/turistikrota/service.notify/src/adapters/mail"
	"github.com/turistikrota/service.notify/src/adapters/mongo"
	"github.com/turistikrota/service.notify/src/adapters/sms"
	"github.com/turistikrota/service.notify/src/adapters/telegram"
)

var (
	Mail     = mail.New()
	Telegram = telegram.NewCreator()
	SMS      = sms.New()
	Mongo    = mongo.New()
)
