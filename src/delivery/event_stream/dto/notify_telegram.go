package dto

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/domain/telegram"
)

type NotifyTelegram struct {
	telegram.Data
	Recipient string `json:"recipient"`
}

func (n *NotifyTelegram) ToCommand() command.SendTelegramCommand {
	return command.SendTelegramCommand{
		Recipient: n.Recipient,
		Data:      &n.Data,
	}
}
