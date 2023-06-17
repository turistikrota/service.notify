package dto

import (
	"api.turistikrota.com/notify/src/app/command"
	"api.turistikrota.com/notify/src/domain/sms"
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
