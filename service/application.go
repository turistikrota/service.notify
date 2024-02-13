package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.notify/adapters/mail"
	"github.com/turistikrota/service.notify/adapters/sms"
	"github.com/turistikrota/service.notify/adapters/telegram"
	"github.com/turistikrota/service.notify/app"
	"github.com/turistikrota/service.notify/app/command"
	"github.com/turistikrota/service.notify/app/query"
	"github.com/turistikrota/service.notify/config"
	"github.com/turistikrota/service.notify/domains/actor_config"
	"github.com/turistikrota/service.notify/domains/notify"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
	I18n        *i18np.I18n
}

func NewApplication(cnf Config) app.Application {

	actorConfigFactory := actor_config.NewFactory()
	actorConfigRepo := actor_config.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.ActorConfig.Collection), actorConfigFactory)

	notifyFactory := notify.NewFactory()

	mail := mail.New(cnf.App.Smtp)
	sms := sms.New(cnf.App.Adapters.NetGsm)
	telegram := telegram.New(cnf.App.Adapters.Telegram)

	return app.Application{
		Commands: app.Commands{
			ActorConfigCreateUser:     command.NewActorConfigCreateUserHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigCreateBusiness: command.NewActorConfigCreateBusinessHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigAddMail:        command.NewActorConfigAddMailHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigAddSms:         command.NewActorConfigAddSmsHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigAddTelegram:    command.NewActorConfigAddTelegramHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigUpdateMail:     command.NewActorConfigUpdateMailHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigUpdateSms:      command.NewActorConfigUpdateSmsHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigUpdateTelegram: command.NewActorConfigUpdateTelegramHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigRemoveMail:     command.NewActorConfigRemoveMailHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigRemoveSms:      command.NewActorConfigRemoveSmsHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigRemoveTelegram: command.NewActorConfigRemoveTelegramHandler(actorConfigFactory, actorConfigRepo),
			NotifyTestMail:            command.NewNotifyTestMailHandler(notifyFactory, cnf.I18n, mail),
			NotifyTestSms:             command.NewNotifyTestSmsHandler(notifyFactory, cnf.I18n, sms),
			NotifyTestTelegram:        command.NewNotifyTestTelegramHandler(notifyFactory, cnf.I18n, telegram),
			NotifySendEmail:           command.NewNotifySendEmailHandler(notifyFactory, actorConfigRepo, cnf.I18n, mail),
			NotifySendSms:             command.NewNotifySendSmsHandler(notifyFactory, actorConfigRepo, cnf.I18n, sms),
			NotifySendSpecialEmail:    command.NewNotifySendSpecialEmailHandler(notifyFactory, cnf.I18n, mail),
			NotifySendSpecialSms:      command.NewNotifySendSpecialSmsHandler(notifyFactory, cnf.I18n, sms),
			NotifySendToAllChannels:   command.NewNotifySendToAllChannelsHandler(notifyFactory, actorConfigRepo, cnf.I18n, sms, mail, telegram),
		},
		Queries: app.Queries{
			ActorConfigFilter:            query.NewActorConfigFilterHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigGetByBusinessUUID: query.NewActorConfigGetByBusinessUUIDHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigGetByBusiness:     query.NewActorConfigGetByBusinessHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigGetByUserName:     query.NewActorConfigGetByUserNameHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigGetByUser:         query.NewActorConfigGetByUserHandler(actorConfigFactory, actorConfigRepo),
			ActorConfigGetByUUID:         query.NewActorConfigGetByUUIDHandler(actorConfigFactory, actorConfigRepo),
		},
	}
}
