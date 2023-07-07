package service

import (
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.notify/src/adapters"
	"github.com/turistikrota/service.notify/src/app"
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/app/query"
	"github.com/turistikrota/service.notify/src/config"
	"github.com/turistikrota/service.notify/src/domain/mail"
	"github.com/turistikrota/service.notify/src/domain/notify"
	"github.com/turistikrota/service.notify/src/domain/sms"
	"github.com/turistikrota/service.notify/src/domain/telegram"
	"github.com/turistikrota/service.shared/db/mongo"
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
