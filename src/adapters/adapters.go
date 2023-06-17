package adapters

import (
	"api.turistikrota.com/notify/src/adapters/mail"
	"api.turistikrota.com/notify/src/adapters/mongo"
	"api.turistikrota.com/notify/src/adapters/sms"
	"api.turistikrota.com/notify/src/adapters/telegram"
)

var (
	Mail     = mail.New()
	Telegram = telegram.NewCreator()
	SMS      = sms.New()
	Mongo    = mongo.New()
)
