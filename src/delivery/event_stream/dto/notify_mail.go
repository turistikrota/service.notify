package dto

import (
	"api.turistikrota.com/notify/src/app/command"
	"api.turistikrota.com/notify/src/domain/mail"
)

type NotifyMail struct {
	mail.Data
	Recipient string `json:"recipient"`
}

func (n *NotifyMail) ToCommand() command.SendMailCommand {
	return command.SendMailCommand{
		Recipient: n.Recipient,
		Data:      &n.Data,
	}
}
