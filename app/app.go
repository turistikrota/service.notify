package app

import (
	"github.com/turistikrota/service.notify/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ActorConfigCreate         command.ActorConfigCreateHandler
	ActorConfigAddMail        command.ActorConfigAddMailHandler
	ActorConfigAddSms         command.ActorConfigAddSmsHandler
	ActorConfigAddTelegram    command.ActorConfigAddTelegramHandler
	ActorConfigUpdateMail     command.ActorConfigUpdateMailHandler
	ActorConfigUpdateSms      command.ActorConfigUpdateSmsHandler
	ActorConfigUpdateTelegram command.ActorConfigUpdateTelegramHandler
	ActorConfigRemoveMail     command.ActorConfigRemoveMailHandler
	ActorConfigRemoveSms      command.ActorConfigRemoveSmsHandler
	ActorConfigRemoveTelegram command.ActorConfigRemoveTelegramHandler
}

type Queries struct {
}
