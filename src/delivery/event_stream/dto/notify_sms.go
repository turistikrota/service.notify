package dto

import (
	"github.com/turistikrota/service.notify/src/app/command"
	"github.com/turistikrota/service.notify/src/domain/sms"
)

type NotifySMS struct {
	sms.Data
	Recipient string `json:"recipient"`
}

func (n *NotifySMS) ToCommand() command.SendSmsCommand {
	return command.SendSmsCommand{
		Recipient: n.Recipient,
		Data:      &n.Data,
	}
}
