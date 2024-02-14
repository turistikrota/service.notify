package app

import (
	"github.com/turistikrota/service.notify/app/command"
	"github.com/turistikrota/service.notify/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ActorConfigCreateUser     command.ActorConfigCreateUserHandler
	ActorConfigCreateBusiness command.ActorConfigCreateBusinessHandler
	ActorConfigAdd            command.ActorConfigAddHandler
	ActorConfigUpdate         command.ActorConfigUpdateHandler
	ActorConfigRemove         command.ActorConfigRemoveHandler

	NotifyTestMail     command.NotifyTestMailHandler
	NotifyTestSms      command.NotifyTestSmsHandler
	NotifyTestTelegram command.NotifyTestTelegramHandler

	NotifySendEmail         command.NotifySendEmailHandler
	NotifySendSms           command.NotifySendSmsHandler
	NotifySendPush          command.NotifySendPushHandler
	NotifySendSpecialEmail  command.NotifySendSpecialEmailHandler
	NotifySendSpecialSms    command.NotifySendSpecialSmsHandler
	NotifySendToAllChannels command.NotifySendToAllChannelsHandler
}

type Queries struct {
	ActorConfigFilter                query.ActorConfigFilterHandler
	ActorConfigGetByBusinessUUID     query.ActorConfigGetByBusinessUUIDHandler
	ActorConfigGetByBusiness         query.ActorConfigGetByBusinessHandler
	ActorConfigGetByUserName         query.ActorConfigGetByUserNameHandler
	ActorConfigGetByUser             query.ActorConfigGetByUserHandler
	ActorConfigGetByUUID             query.ActorConfigGetByUUIDHandler
	ActorConfigGetOrCreateByBusiness query.ActorConfigGetOrCreateByBusinessHandler
	ActorConfigGetOrCreateByUser     query.ActorConfigGetOrCreateByUserHandler
}
