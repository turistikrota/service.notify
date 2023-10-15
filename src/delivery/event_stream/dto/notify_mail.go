package dto

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/domain/mail"
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
