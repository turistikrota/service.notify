package service

import (
	"api.turistikrota.com/notify/src/adapters"
	"api.turistikrota.com/notify/src/app"
	"api.turistikrota.com/notify/src/app/command"
	"api.turistikrota.com/notify/src/app/query"
	"api.turistikrota.com/notify/src/config"
	"api.turistikrota.com/notify/src/domain/mail"
	"api.turistikrota.com/notify/src/domain/notify"
	"api.turistikrota.com/notify/src/domain/sms"
	"api.turistikrota.com/notify/src/domain/telegram"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/decorator"
	"github.com/turistikrota/service.shared/events"
	"github.com/turistikrota/service.shared/validator"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
}

func NewApplication(c Config) app.Application {
	notifyFactory := notify.NewFactory(c.Validator, false)
	notifyRepo := adapters.Mongo.NewNotify(notifyFactory, c.Mongo.GetCollection(c.App.DB.MongoNotify.Collection))

	mailFactory := mail.NewFactory(c.Validator, notifyFactory)
	mailRepo := adapters.Mail.NewGoogle(mailFactory, notifyFactory, c.App.Adapters.MailGoogle)

	smsFactory := sms.NewFactory(c.Validator, notifyFactory)
	smsRepo := adapters.SMS.NewNetGSM(smsFactory, notifyFactory, c.App.Adapters.NetGsm)

	telegramFactory := telegram.NewFactory(c.Validator, notifyFactory)
	telegramRepo := adapters.Telegram.New(telegramFactory, notifyFactory, c.App.Adapters.Telegram)

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			SendMail: command.NewSendMailHandler(command.SendMailHandlerConfig{
				NotifyRepo:    notifyRepo,
				MailRepo:      mailRepo,
				MailFactory:   mailFactory,
				NotifyFactory: notifyFactory,
				CqrsBase:      base,
			}),
			SendSms: command.NewSendSmsHandler(command.SendSmsHandlerConfig{
				NotifyRepo:    notifyRepo,
				SmsRepo:       smsRepo,
				SmsFactory:    smsFactory,
				NotifyFactory: notifyFactory,
				CqrsBase:      base,
			}),
			SendTelegram: command.NewSendTelegramHandler(command.SendTelegramHandlerConfig{
				NotifyRepo:      notifyRepo,
				TelegramRepo:    telegramRepo,
				TelegramFactory: telegramFactory,
				NotifyFactory:   notifyFactory,
				CqrsBase:        base,
			}),
		},
		Queries: app.Queries{
			GetByUUID: query.NewGetByUUIDHandler(query.GetByUUIDHandlerConfig{
				NotifyRepo: notifyRepo,
				CqrsBase:   base,
			}),
			GetAllByChannel: query.NewGetAllByChannelHandler(query.GetAllByChannelHandlerConfig{
				NotifyRepo: notifyRepo,
				CqrsBase:   base,
			}),
			GetAllByRecipient: query.NewGetAllByRecipientHandler(query.GetAllByRecipientHandlerConfig{
				NotifyRepo: notifyRepo,
				CqrsBase:   base,
			}),
		},
	}
}
