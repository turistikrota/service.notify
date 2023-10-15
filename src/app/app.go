package app

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendMail     command.SendMailHandler
	SendSms      command.SendSmsHandler
	SendTelegram command.SendTelegramHandler
	SendPush     command.SendPushHandler
}

type Queries struct {
	GetByUUID         query.GetByUUIDHandler
	GetAllByChannel   query.GetAllByChannelHandler
	GetAllByRecipient query.GetAllByRecipientHandler
}
