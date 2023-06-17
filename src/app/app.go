package app

import (
	"api.turistikrota.com/notify/src/app/command"
	"api.turistikrota.com/notify/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendMail     command.SendMailHandler
	SendSms      command.SendSmsHandler
	SendTelegram command.SendTelegramHandler
}

type Queries struct {
	GetByUUID         query.GetByUUIDHandler
	GetAllByChannel   query.GetAllByChannelHandler
	GetAllByRecipient query.GetAllByRecipientHandler
}
