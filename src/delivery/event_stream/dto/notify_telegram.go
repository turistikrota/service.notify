package dto

import (
	"api.turistikrota.com/notify/src/app/command"
	"api.turistikrota.com/notify/src/domain/telegram"
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
